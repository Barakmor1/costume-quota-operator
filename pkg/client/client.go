package client

//go:generate mockgen -source $GOFILE -package=$GOPACKAGE -destination=generated_mock_$GOFILE

/*
 ATTENTION: Rerun code generators when interface signatures are modified.
*/

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	generatedclient "kubevirt.io/application-aware-quota/pkg/generated/aaq/clientset/versioned"
	aaqv1alpha1 "kubevirt.io/application-aware-quota/pkg/generated/aaq/clientset/versioned/typed/core/v1alpha1"
	crqclient "kubevirt.io/application-aware-quota/pkg/generated/cluster-resource-quota/clientset/versioned"
	kubevirtclient "kubevirt.io/application-aware-quota/pkg/generated/kubevirt/clientset/versioned"
	"kubevirt.io/application-aware-quota/staging/src/kubevirt.io/applications-aware-quota-api/pkg/apis/core/v1alpha1"
)

type AAQClient interface {
	RestClient() *rest.RESTClient
	kubernetes.Interface
	ApplicationAwareResourceQuotas(namespace string) ApplicationAwareResourceQuotaInterface
	ApplicationAwareClusterResourceQuotas() ApplicationAwareClusterResourceQuotaInterface
	AAQJobQueueConfigs(namespace string) AAQJobQueueConfigInterface
	AAQ() AAQInterface
	GeneratedAAQClient() generatedclient.Interface
	CRQClient() crqclient.Interface
	KubevirtClient() kubevirtclient.Interface
	DiscoveryClient() discovery.DiscoveryInterface
	Config() *rest.Config
}

type aaq struct {
	master             string
	kubeconfig         string
	restClient         *rest.RESTClient
	config             *rest.Config
	generatedAAQClient *generatedclient.Clientset
	kubevirtClient     *kubevirtclient.Clientset
	crqClient          *crqclient.Clientset
	discoveryClient    *discovery.DiscoveryClient
	dynamicClient      dynamic.Interface
	*kubernetes.Clientset
}

func (k aaq) CRQClient() crqclient.Interface {
	return k.crqClient
}

func (k aaq) KubevirtClient() kubevirtclient.Interface {
	return k.kubevirtClient
}

func (k aaq) Config() *rest.Config {
	return k.config
}

func (k aaq) RestClient() *rest.RESTClient {
	return k.restClient
}

func (k aaq) GeneratedAAQClient() generatedclient.Interface {
	return k.generatedAAQClient
}

func (k aaq) ApplicationAwareResourceQuotas(namespace string) ApplicationAwareResourceQuotaInterface {
	return k.generatedAAQClient.AaqV1alpha1().ApplicationAwareResourceQuotas(namespace)
}

func (k aaq) ApplicationAwareClusterResourceQuotas() ApplicationAwareClusterResourceQuotaInterface {
	return k.generatedAAQClient.AaqV1alpha1().ApplicationAwareClusterResourceQuotas()
}

func (k aaq) AAQJobQueueConfigs(namespace string) AAQJobQueueConfigInterface {
	return k.generatedAAQClient.AaqV1alpha1().AAQJobQueueConfigs(namespace)
}
func (k aaq) AAQ() AAQInterface {
	return k.generatedAAQClient.AaqV1alpha1().AAQs()
}

func (k aaq) DynamicClient() dynamic.Interface {
	return k.dynamicClient
}

func (k aaq) DiscoveryClient() discovery.DiscoveryInterface {
	return k.discoveryClient
}

// ApplicationAwareResourceQuotaInterface has methods to work with ApplicationAwareResourceQuotas resources.
type ApplicationAwareResourceQuotaInterface interface {
	Create(ctx context.Context, applicationAwareResourceQuota *v1alpha1.ApplicationAwareResourceQuota, opts metav1.CreateOptions) (*v1alpha1.ApplicationAwareResourceQuota, error)
	Update(ctx context.Context, applicationAwareResourceQuota *v1alpha1.ApplicationAwareResourceQuota, opts metav1.UpdateOptions) (*v1alpha1.ApplicationAwareResourceQuota, error)
	UpdateStatus(ctx context.Context, applicationAwareResourceQuota *v1alpha1.ApplicationAwareResourceQuota, opts metav1.UpdateOptions) (*v1alpha1.ApplicationAwareResourceQuota, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.ApplicationAwareResourceQuota, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.ApplicationAwareResourceQuotaList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationAwareResourceQuota, err error)
	aaqv1alpha1.ApplicationAwareResourceQuotaExpansion
}

// ApplicationAwareClusterResourceQuotaInterface has methods to work with ApplicationAwareClusterResourceQuota resources.
type ApplicationAwareClusterResourceQuotaInterface interface {
	Create(ctx context.Context, applicationAwareResourceQuota *v1alpha1.ApplicationAwareClusterResourceQuota, opts metav1.CreateOptions) (*v1alpha1.ApplicationAwareClusterResourceQuota, error)
	Update(ctx context.Context, applicationAwareResourceQuota *v1alpha1.ApplicationAwareClusterResourceQuota, opts metav1.UpdateOptions) (*v1alpha1.ApplicationAwareClusterResourceQuota, error)
	UpdateStatus(ctx context.Context, applicationAwareResourceQuota *v1alpha1.ApplicationAwareClusterResourceQuota, opts metav1.UpdateOptions) (*v1alpha1.ApplicationAwareClusterResourceQuota, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.ApplicationAwareClusterResourceQuota, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.ApplicationAwareClusterResourceQuotaList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationAwareClusterResourceQuota, err error)
	aaqv1alpha1.ApplicationAwareClusterResourceQuotaExpansion
}

// AAQJobQueueConfigInterface has methods to work with AAQJobQueueConfigs resources.
type AAQJobQueueConfigInterface interface {
	Create(ctx context.Context, aAQJobQueueConfig *v1alpha1.AAQJobQueueConfig, opts metav1.CreateOptions) (*v1alpha1.AAQJobQueueConfig, error)
	Update(ctx context.Context, aAQJobQueueConfig *v1alpha1.AAQJobQueueConfig, opts metav1.UpdateOptions) (*v1alpha1.AAQJobQueueConfig, error)
	UpdateStatus(ctx context.Context, aAQJobQueueConfig *v1alpha1.AAQJobQueueConfig, opts metav1.UpdateOptions) (*v1alpha1.AAQJobQueueConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.AAQJobQueueConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.AAQJobQueueConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.AAQJobQueueConfig, err error)
	aaqv1alpha1.AAQJobQueueConfigExpansion
}

// AAQInterface has methods to work with AAQ resources.
type AAQInterface interface {
	Create(ctx context.Context, aAQ *v1alpha1.AAQ, opts metav1.CreateOptions) (*v1alpha1.AAQ, error)
	Update(ctx context.Context, aAQ *v1alpha1.AAQ, opts metav1.UpdateOptions) (*v1alpha1.AAQ, error)
	UpdateStatus(ctx context.Context, aAQ *v1alpha1.AAQ, opts metav1.UpdateOptions) (*v1alpha1.AAQ, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.AAQ, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.AAQList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.AAQ, err error)
	aaqv1alpha1.AAQExpansion
}
