package arq_controller

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	k8sadmission "k8s.io/apiserver/pkg/admission"
	quotaplugin "k8s.io/apiserver/pkg/admission/plugin/resourcequota"
	"k8s.io/apiserver/pkg/admission/plugin/resourcequota/apis/resourcequota"
	v12 "k8s.io/apiserver/pkg/quota/v1"
	"k8s.io/client-go/kubernetes/scheme"
	v14 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
	_ "kubevirt.io/api/core/v1"
	aaq_evaluator "kubevirt.io/applications-aware-quota/pkg/aaq-controller/aaq-evaluator"
	built_in_usage_calculators "kubevirt.io/applications-aware-quota/pkg/aaq-controller/built-in-usage-calculators"
	"kubevirt.io/applications-aware-quota/pkg/aaq-operator/resources/utils"
	v1alpha13 "kubevirt.io/applications-aware-quota/pkg/generated/clientset/versioned/typed/core/v1alpha1"
	v1alpha12 "kubevirt.io/applications-aware-quota/staging/src/kubevirt.io/applications-aware-quota-api/pkg/apis/core/v1alpha1"
	"kubevirt.io/client-go/kubecli"
	"kubevirt.io/client-go/log"
	corev1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
	"time"
)

type enqueueState string

const (
	Immediate  enqueueState = "Immediate"
	Forget     enqueueState = "Forget"
	BackOff    enqueueState = "BackOff"
	AaqjqcName string       = "aaqjqc"
)

type AaqGateController struct {
	podInformer    cache.SharedIndexInformer
	arqInformer    cache.SharedIndexInformer
	aaqjqcInformer cache.SharedIndexInformer
	arqQueue       workqueue.RateLimitingInterface
	clientSet      kubecli.KubevirtClient
	aaqCli         v1alpha13.AaqV1alpha1Client
	recorder       record.EventRecorder
	aaqEvaluator   v12.Evaluator
}

func NewAaqGateController(clientSet kubecli.KubevirtClient,
	aaqCli v1alpha13.AaqV1alpha1Client,
	podInformer cache.SharedIndexInformer,
	arqInformer cache.SharedIndexInformer,
	aaqjqcInformer cache.SharedIndexInformer,
	stop <-chan struct{},
) *AaqGateController {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&v14.EventSinkImpl{Interface: clientSet.CoreV1().Events(v1.NamespaceAll)})

	//todo: make this generic for now we will try only launcher calculator
	calcRegistry := aaq_evaluator.NewAaqCalculatorsRegistry(10, clock.RealClock{}).AddCalculator(built_in_usage_calculators.NewVirtLauncherCalculator(stop))

	ctrl := AaqGateController{
		clientSet:      clientSet,
		aaqCli:         aaqCli,
		aaqjqcInformer: aaqjqcInformer,
		podInformer:    podInformer,
		arqInformer:    arqInformer,
		arqQueue:       workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "arq-queue"),
		recorder:       eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: utils.ControllerPodName}),
		aaqEvaluator:   aaq_evaluator.NewAaqEvaluator(nil, calcRegistry, clock.RealClock{}),
	}

	return &ctrl
}

func (ctrl *AaqGateController) addAllArqsInNamespace(ns string) {
	objs, err := ctrl.arqInformer.GetIndexer().ByIndex(cache.NamespaceIndex, ns)
	atLeastOneArq := false
	if err != nil {
		log.Log.Infof("AaqGateController: Error failed to list pod from podInformer")
	}
	for _, obj := range objs {
		arq := obj.(*v1alpha12.ApplicationsResourceQuota)
		key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(arq)
		if err != nil {
			return
		}
		atLeastOneArq = true
		ctrl.arqQueue.Add(key)
	}
	if !atLeastOneArq {
		ctrl.arqQueue.Add(ns + "/fake")
	}
}

// When a ApplicationsResourceQuota is deleted, enqueue all gated pods for revaluation
func (ctrl *AaqGateController) deleteArq(obj interface{}) {
	arq := obj.(*v1alpha12.ApplicationsResourceQuota)
	ctrl.addAllArqsInNamespace(arq.Namespace)
	return
}

// When a ApplicationsResourceQuota is updated, enqueue all gated pods for revaluation
func (ctrl *AaqGateController) addArq(obj interface{}) {
	arq := obj.(*v1alpha12.ApplicationsResourceQuota)
	ctrl.addAllArqsInNamespace(arq.Namespace)
	return
}

// When a ApplicationsResourceQuota is updated, enqueue all gated pods for revaluation
func (ctrl *AaqGateController) updateArq(old, cur interface{}) {
	arq := cur.(*v1alpha12.ApplicationsResourceQuota)
	ctrl.addAllArqsInNamespace(arq.Namespace)
	return
}

// When a ApplicationsResourceQuotaaqjqc.Status.PodsInJobQueuea is updated, enqueue all gated pods for revaluation
func (ctrl *AaqGateController) updateAaqjqc(old, cur interface{}) {
	aaqjqc := cur.(*v1alpha12.AAQJobQueueConfig)
	if aaqjqc.Status.PodsInJobQueue == nil || len(aaqjqc.Status.PodsInJobQueue) == 0 {
		ctrl.addAllArqsInNamespace(aaqjqc.Namespace)

	}
	return
}

// When a ApplicationsResourceQuota is updated, enqueue all gated pods for revaluation
func (ctrl *AaqGateController) deleteAaqjqc(obj interface{}) {
	arq := obj.(*v1alpha12.AAQJobQueueConfig)
	ctrl.addAllArqsInNamespace(arq.Namespace)
	return
}

func (ctrl *AaqGateController) addPod(obj interface{}) {
	pod := obj.(*v1.Pod)
	log.Log.Infof(fmt.Sprintf("Barak: here %v", pod))

	if pod.Spec.SchedulingGates != nil &&
		len(pod.Spec.SchedulingGates) == 1 &&
		pod.Spec.SchedulingGates[0].Name == "ApplicationsAwareQuotaGate" {
		ctrl.addAllArqsInNamespace(pod.Namespace)
	}
}
func (ctrl *AaqGateController) updatePod(old, curr interface{}) {
	pod := curr.(*v1.Pod)
	log.Log.Infof(fmt.Sprintf("Barak: here %v", pod))

	if pod.Spec.SchedulingGates != nil &&
		len(pod.Spec.SchedulingGates) == 1 &&
		pod.Spec.SchedulingGates[0].Name == "ApplicationsAwareQuotaGate" {
		ctrl.addAllArqsInNamespace(pod.Namespace)
	}
}

func (ctrl *AaqGateController) runWorker() {
	for ctrl.Execute() {
	}
}

func (ctrl *AaqGateController) Execute() bool {
	log.Log.Infof("Barak: here")

	key, quit := ctrl.arqQueue.Get()
	if quit {
		return false
	}
	defer ctrl.arqQueue.Done(key)

	err, enqueueState := ctrl.execute(key.(string))
	if err != nil {
		klog.Errorf(fmt.Sprintf("AaqGateController: Error with key: %v err: %v", key, err))
	}
	switch enqueueState {
	case BackOff:
		ctrl.arqQueue.AddRateLimited(key)
	case Forget:
		ctrl.arqQueue.Forget(key)
	case Immediate:
		ctrl.arqQueue.Add(key)
	}

	return true
}

func (ctrl *AaqGateController) execute(key string) (error, enqueueState) {

	arqNS, _, err := cache.SplitMetaNamespaceKey(key)
	aaqjqc, err := ctrl.aaqCli.AAQJobQueueConfigs(arqNS).Get(context.Background(), AaqjqcName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		aaqjqc, err = ctrl.aaqCli.AAQJobQueueConfigs(arqNS).Create(context.Background(),
			&v1alpha12.AAQJobQueueConfig{ObjectMeta: metav1.ObjectMeta{Name: AaqjqcName}, Spec: v1alpha12.AAQJobQueueConfigSpec{}},
			metav1.CreateOptions{})
		if err != nil {
			log.Log.Infof(fmt.Sprintf("Barak: here %v", err.Error()))
			return err, Immediate
		}
	} else if err != nil {
		log.Log.Infof(fmt.Sprintf("Barak: here %v", err.Error()))
		return err, Immediate
	}
	if len(aaqjqc.Status.PodsInJobQueue) != 0 {
		ctrl.releasePods(aaqjqc.Status.PodsInJobQueue, arqNS)
		return nil, Immediate //wait for the calculator to process changes
	}

	arqsObjs, err := ctrl.arqInformer.GetIndexer().ByIndex(cache.NamespaceIndex, arqNS)
	if err != nil {
		return err, Immediate
	}
	var rqs []v1.ResourceQuota
	for _, arqObj := range arqsObjs {
		arq := arqObj.(*v1alpha12.ApplicationsResourceQuota)
		rqs = append(rqs, v1.ResourceQuota{ObjectMeta: metav1.ObjectMeta{Name: arq.Name, Namespace: arqNS}, Spec: arq.Spec, Status: arq.Status})
	}
	podObjs, err := ctrl.podInformer.GetIndexer().ByIndex(cache.NamespaceIndex, arqNS)
	if err != nil {
		log.Log.Infof(fmt.Sprintf("Barak: here %v", err.Error()))
		return err, Immediate
	}
	for _, podObj := range podObjs {
		pod := podObj.(*v1.Pod)
		if pod.Spec.SchedulingGates != nil &&
			len(pod.Spec.SchedulingGates) == 1 &&
			pod.Spec.SchedulingGates[0].Name == "ApplicationsAwareQuotaGate" {
			podCopy := pod.DeepCopy()
			if podCopy.Spec.SchedulingGates != nil && len(podCopy.Spec.SchedulingGates) == 1 && podCopy.Spec.SchedulingGates[0].Name == "ApplicationsAwareQuotaGate" {
				podCopy.Spec.SchedulingGates = []v1.PodSchedulingGate{}
			}
			podToCreateAttr := k8sadmission.NewAttributesRecord(podCopy, nil,
				apiextensions.Kind("Pod").WithVersion("version"), podCopy.Namespace, podCopy.Name,
				corev1beta1.Resource("pods").WithVersion("version"), "", k8sadmission.Create,
				&metav1.CreateOptions{}, false, nil)

			currPodLimitedResource, err := getCurrLimitedResource(ctrl.aaqEvaluator, podCopy)

			if err != nil {
				log.Log.Infof(fmt.Sprintf("Barak: here %v", err.Error()))
				return nil, Immediate
			}

			newRq, err := quotaplugin.CheckRequest(rqs, podToCreateAttr, ctrl.aaqEvaluator, []resourcequota.LimitedResource{currPodLimitedResource})
			log.Log.Infof("barak here1")
			if err == nil {
				rqs = newRq
				aaqjqc.Status.PodsInJobQueue = append(aaqjqc.Status.PodsInJobQueue, pod.Name)
			} //todo: create an event if we are blocked for a while
		}
	}

	if len(aaqjqc.Status.PodsInJobQueue) > 0 {
		aaqjqc, err = ctrl.aaqCli.AAQJobQueueConfigs(arqNS).UpdateStatus(context.Background(), aaqjqc, metav1.UpdateOptions{})
		if err != nil {
			log.Log.Infof(fmt.Sprintf("Barak: here %v", err.Error()))
			return err, Immediate
		}
	}
	err = ctrl.releasePods(aaqjqc.Status.PodsInJobQueue, arqNS)
	if err != nil {
		log.Log.Infof(fmt.Sprintf("Barak: here %v", err.Error()))
		return err, Immediate
	}
	return nil, Forget
}

func (ctrl *AaqGateController) releasePods(podsToRelease []string, ns string) error {
	for _, podName := range podsToRelease {
		obj, exists, err := ctrl.podInformer.GetIndexer().GetByKey(ns + "/" + podName)
		if err != nil {
			return err
		}
		if !exists {
			continue
		}
		pod := obj.(*v1.Pod)
		if pod.Spec.SchedulingGates != nil && len(pod.Spec.SchedulingGates) == 1 && pod.Spec.SchedulingGates[0].Name == "ApplicationsAwareQuotaGate" {
			pod.Spec.SchedulingGates = []v1.PodSchedulingGate{}
			_, err = ctrl.clientSet.CoreV1().Pods(ns).Update(context.Background(), pod, metav1.UpdateOptions{})
			if err != nil {
				return err
			}
		}
	}
	return nil

}

func (ctrl *AaqGateController) Run(threadiness int, stop <-chan struct{}) error {
	defer utilruntime.HandleCrash()

	_, err := ctrl.podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    ctrl.addPod,
		UpdateFunc: ctrl.updatePod,
	})
	if err != nil {
		return err
	}
	_, err = ctrl.arqInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: ctrl.deleteArq,
		UpdateFunc: ctrl.updateArq,
		AddFunc:    ctrl.addArq,
	})
	if err != nil {
		return err
	}
	_, err = ctrl.aaqjqcInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: ctrl.deleteAaqjqc,
		UpdateFunc: ctrl.updateAaqjqc,
	})
	if err != nil {
		return err
	}
	klog.Info("Starting Arq controller")
	defer klog.Info("Shutting down Arq controller")

	for i := 0; i < threadiness; i++ {
		go wait.Until(ctrl.runWorker, time.Second, stop)
	}

	<-stop
	return nil

}

func getCurrLimitedResource(podEvaluator v12.Evaluator, podToCreate *v1.Pod) (resourcequota.LimitedResource, error) {
	launcherLimitedResource := resourcequota.LimitedResource{
		Resource:      "pods",
		MatchContains: []string{},
	}
	usage, err := podEvaluator.Usage(podToCreate)
	if err != nil {
		return launcherLimitedResource, err
	}
	for k, _ := range usage {
		launcherLimitedResource.MatchContains = append(launcherLimitedResource.MatchContains, string(k))
	}
	return launcherLimitedResource, nil
}
