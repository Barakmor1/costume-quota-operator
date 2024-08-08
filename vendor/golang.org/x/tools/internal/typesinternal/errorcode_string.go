// Code generated by "stringer -type=ErrorCode"; DO NOT EDIT.

package typesinternal

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidSyntaxTree - -1]
	_ = x[Test-1]
	_ = x[BlankPkgName-2]
	_ = x[MismatchedPkgName-3]
	_ = x[InvalidPkgUse-4]
	_ = x[BadImportPath-5]
	_ = x[BrokenImport-6]
	_ = x[ImportCRenamed-7]
	_ = x[UnusedImport-8]
	_ = x[InvalidInitCycle-9]
	_ = x[DuplicateDecl-10]
	_ = x[InvalidDeclCycle-11]
	_ = x[InvalidTypeCycle-12]
	_ = x[InvalidConstInit-13]
	_ = x[InvalidConstVal-14]
	_ = x[InvalidConstType-15]
	_ = x[UntypedNilUse-16]
	_ = x[WrongAssignCount-17]
	_ = x[UnassignableOperand-18]
	_ = x[NoNewVar-19]
	_ = x[MultiValAssignOp-20]
	_ = x[InvalidIfaceAssign-21]
	_ = x[InvalidChanAssign-22]
	_ = x[IncompatibleAssign-23]
	_ = x[UnaddressableFieldAssign-24]
	_ = x[NotAType-25]
	_ = x[InvalidArrayLen-26]
	_ = x[BlankIfaceMethod-27]
	_ = x[IncomparableMapKey-28]
	_ = x[InvalidIfaceEmbed-29]
	_ = x[InvalidPtrEmbed-30]
	_ = x[BadRecv-31]
	_ = x[InvalidRecv-32]
	_ = x[DuplicateFieldAndMethod-33]
	_ = x[DuplicateMethod-34]
	_ = x[InvalidBlank-35]
	_ = x[InvalidIota-36]
	_ = x[MissingInitBody-37]
	_ = x[InvalidInitSig-38]
	_ = x[InvalidInitDecl-39]
	_ = x[InvalidMainDecl-40]
	_ = x[TooManyValues-41]
	_ = x[NotAnExpr-42]
	_ = x[TruncatedFloat-43]
	_ = x[NumericOverflow-44]
	_ = x[UndefinedOp-45]
	_ = x[MismatchedTypes-46]
	_ = x[DivByZero-47]
	_ = x[NonNumericIncDec-48]
	_ = x[UnaddressableOperand-49]
	_ = x[InvalidIndirection-50]
	_ = x[NonIndexableOperand-51]
	_ = x[InvalidIndex-52]
	_ = x[SwappedSliceIndices-53]
	_ = x[NonSliceableOperand-54]
	_ = x[InvalidSliceExpr-55]
	_ = x[InvalidShiftCount-56]
	_ = x[InvalidShiftOperand-57]
	_ = x[InvalidReceive-58]
	_ = x[InvalidSend-59]
	_ = x[DuplicateLitKey-60]
	_ = x[MissingLitKey-61]
	_ = x[InvalidLitIndex-62]
	_ = x[OversizeArrayLit-63]
	_ = x[MixedStructLit-64]
	_ = x[InvalidStructLit-65]
	_ = x[MissingLitField-66]
	_ = x[DuplicateLitField-67]
	_ = x[UnexportedLitField-68]
	_ = x[InvalidLitField-69]
	_ = x[UntypedLit-70]
	_ = x[InvalidLit-71]
	_ = x[AmbiguousSelector-72]
	_ = x[UndeclaredImportedName-73]
	_ = x[UnexportedName-74]
	_ = x[UndeclaredName-75]
	_ = x[MissingFieldOrMethod-76]
	_ = x[BadDotDotDotSyntax-77]
	_ = x[NonVariadicDotDotDot-78]
	_ = x[MisplacedDotDotDot-79]
	_ = x[InvalidDotDotDotOperand-80]
	_ = x[InvalidDotDotDot-81]
	_ = x[UncalledBuiltin-82]
	_ = x[InvalidAppend-83]
	_ = x[InvalidCap-84]
	_ = x[InvalidClose-85]
	_ = x[InvalidCopy-86]
	_ = x[InvalidComplex-87]
	_ = x[InvalidDelete-88]
	_ = x[InvalidImag-89]
	_ = x[InvalidLen-90]
	_ = x[SwappedMakeArgs-91]
	_ = x[InvalidMake-92]
	_ = x[InvalidReal-93]
	_ = x[InvalidAssert-94]
	_ = x[ImpossibleAssert-95]
	_ = x[InvalidConversion-96]
	_ = x[InvalidUntypedConversion-97]
	_ = x[BadOffsetofSyntax-98]
	_ = x[InvalidOffsetof-99]
	_ = x[UnusedExpr-100]
	_ = x[UnusedVar-101]
	_ = x[MissingReturn-102]
	_ = x[WrongResultCount-103]
	_ = x[OutOfScopeResult-104]
	_ = x[InvalidCond-105]
	_ = x[InvalidPostDecl-106]
	_ = x[InvalidChanRange-107]
	_ = x[InvalidIterVar-108]
	_ = x[InvalidRangeExpr-109]
	_ = x[MisplacedBreak-110]
	_ = x[MisplacedContinue-111]
	_ = x[MisplacedFallthrough-112]
	_ = x[DuplicateCase-113]
	_ = x[DuplicateDefault-114]
	_ = x[BadTypeKeyword-115]
	_ = x[InvalidTypeSwitch-116]
	_ = x[InvalidExprSwitch-117]
	_ = x[InvalidSelectCase-118]
	_ = x[UndeclaredLabel-119]
	_ = x[DuplicateLabel-120]
	_ = x[MisplacedLabel-121]
	_ = x[UnusedLabel-122]
	_ = x[JumpOverDecl-123]
	_ = x[JumpIntoBlock-124]
	_ = x[InvalidMethodExpr-125]
	_ = x[WrongArgCount-126]
	_ = x[InvalidCall-127]
	_ = x[UnusedResults-128]
	_ = x[InvalidDefer-129]
	_ = x[InvalidGo-130]
	_ = x[BadDecl-131]
	_ = x[RepeatedDecl-132]
	_ = x[InvalidUnsafeAdd-133]
	_ = x[InvalidUnsafeSlice-134]
	_ = x[UnsupportedFeature-135]
	_ = x[NotAGenericType-136]
	_ = x[WrongTypeArgCount-137]
	_ = x[CannotInferTypeArgs-138]
	_ = x[InvalidTypeArg-139]
	_ = x[InvalidInstanceCycle-140]
	_ = x[InvalidUnion-141]
	_ = x[MisplacedConstraintIface-142]
	_ = x[InvalidMethodTypeParams-143]
	_ = x[MisplacedTypeParam-144]
	_ = x[InvalidUnsafeSliceData-145]
	_ = x[InvalidUnsafeString-146]
}

const (
	_ErrorCode_name_0 = "InvalidSyntaxTree"
	_ErrorCode_name_1 = "TestBlankPkgNameMismatchedPkgNameInvalidPkgUseBadImportPathBrokenImportImportCRenamedUnusedImportInvalidInitCycleDuplicateDeclInvalidDeclCycleInvalidTypeCycleInvalidConstInitInvalidConstValInvalidConstTypeUntypedNilUseWrongAssignCountUnassignableOperandNoNewVarMultiValAssignOpInvalidIfaceAssignInvalidChanAssignIncompatibleAssignUnaddressableFieldAssignNotATypeInvalidArrayLenBlankIfaceMethodIncomparableMapKeyInvalidIfaceEmbedInvalidPtrEmbedBadRecvInvalidRecvDuplicateFieldAndMethodDuplicateMethodInvalidBlankInvalidIotaMissingInitBodyInvalidInitSigInvalidInitDeclInvalidMainDeclTooManyValuesNotAnExprTruncatedFloatNumericOverflowUndefinedOpMismatchedTypesDivByZeroNonNumericIncDecUnaddressableOperandInvalidIndirectionNonIndexableOperandInvalidIndexSwappedSliceIndicesNonSliceableOperandInvalidSliceExprInvalidShiftCountInvalidShiftOperandInvalidReceiveInvalidSendDuplicateLitKeyMissingLitKeyInvalidLitIndexOversizeArrayLitMixedStructLitInvalidStructLitMissingLitFieldDuplicateLitFieldUnexportedLitFieldInvalidLitFieldUntypedLitInvalidLitAmbiguousSelectorUndeclaredImportedNameUnexportedNameUndeclaredNameMissingFieldOrMethodBadDotDotDotSyntaxNonVariadicDotDotDotMisplacedDotDotDotInvalidDotDotDotOperandInvalidDotDotDotUncalledBuiltinInvalidAppendInvalidCapInvalidCloseInvalidCopyInvalidComplexInvalidDeleteInvalidImagInvalidLenSwappedMakeArgsInvalidMakeInvalidRealInvalidAssertImpossibleAssertInvalidConversionInvalidUntypedConversionBadOffsetofSyntaxInvalidOffsetofUnusedExprUnusedVarMissingReturnWrongResultCountOutOfScopeResultInvalidCondInvalidPostDeclInvalidChanRangeInvalidIterVarInvalidRangeExprMisplacedBreakMisplacedContinueMisplacedFallthroughDuplicateCaseDuplicateDefaultBadTypeKeywordInvalidTypeSwitchInvalidExprSwitchInvalidSelectCaseUndeclaredLabelDuplicateLabelMisplacedLabelUnusedLabelJumpOverDeclJumpIntoBlockInvalidMethodExprWrongArgCountInvalidCallUnusedResultsInvalidDeferInvalidGoBadDeclRepeatedDeclInvalidUnsafeAddInvalidUnsafeSliceUnsupportedFeatureNotAGenericTypeWrongTypeArgCountCannotInferTypeArgsInvalidTypeArgInvalidInstanceCycleInvalidUnionMisplacedConstraintIfaceInvalidMethodTypeParamsMisplacedTypeParamInvalidUnsafeSliceDataInvalidUnsafeString"
)

var (
	_ErrorCode_index_1 = [...]uint16{0, 4, 16, 33, 46, 59, 71, 85, 97, 113, 126, 142, 158, 174, 189, 205, 218, 234, 253, 261, 277, 295, 312, 330, 354, 362, 377, 393, 411, 428, 443, 450, 461, 484, 499, 511, 522, 537, 551, 566, 581, 594, 603, 617, 632, 643, 658, 667, 683, 703, 721, 740, 752, 771, 790, 806, 823, 842, 856, 867, 882, 895, 910, 926, 940, 956, 971, 988, 1006, 1021, 1031, 1041, 1058, 1080, 1094, 1108, 1128, 1146, 1166, 1184, 1207, 1223, 1238, 1251, 1261, 1273, 1284, 1298, 1311, 1322, 1332, 1347, 1358, 1369, 1382, 1398, 1415, 1439, 1456, 1471, 1481, 1490, 1503, 1519, 1535, 1546, 1561, 1577, 1591, 1607, 1621, 1638, 1658, 1671, 1687, 1701, 1718, 1735, 1752, 1767, 1781, 1795, 1806, 1818, 1831, 1848, 1861, 1872, 1885, 1897, 1906, 1913, 1925, 1941, 1959, 1977, 1992, 2009, 2028, 2042, 2062, 2074, 2098, 2121, 2139, 2161, 2180}
)

func (i ErrorCode) String() string {
	switch {
	case i == -1:
		return _ErrorCode_name_0
	case 1 <= i && i <= 146:
		i -= 1
		return _ErrorCode_name_1[_ErrorCode_index_1[i]:_ErrorCode_index_1[i+1]]
	default:
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}