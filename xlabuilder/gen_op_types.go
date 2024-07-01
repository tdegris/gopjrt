package xlabuilder

/***** File generated by ./cmd/xlabuilder_codegen, based on op_types.txt. Don't edit it directly. *****/

// OpType enumerates the various operation types supported by XLA.
type OpType int32

const (
	InvalidOp OpType = iota
	ParameterOp
	IotaOp
	ConstantOp
	IdentityOp
	ConvertDTypeOp
	WhereOp
	TupleOp
	GetTupleElementOp
	ReshapeOp
	BroadcastOp
	BroadcastInDimOp
	TransposeOp
	CallOp
	ReduceOp
	ReduceWindowOp
	ConcatenateOp
	SliceOp
	ArgMinMaxOp
	PadOp
	GatherOp
	ScatterOp
	SelectAndScatterOp
	ConvGeneralDilatedOp
	ReverseOp
	DotGeneralOp
	FftOp
	BatchNormTrainingOp
	BatchNormInferenceOp
	BatchNormGradOp
	AbsOp
	NegOp
	ExpOp
	Expm1Op
	FloorOp
	CeilOp
	RoundOp
	LogOp
	Log1pOp
	LogicalNotOp
	LogisticOp
	SignOp
	ClzOp
	CosOp
	SinOp
	TanhOp
	SqrtOp
	RsqrtOp
	ImagOp
	RealOp
	ConjOp
	AddOp
	MulOp
	SubOp
	DivOp
	RemOp
	AndOp
	OrOp
	XorOp
	DotOp
	MinOp
	MaxOp
	PowOp
	ComplexOp
	EqualOp
	NotEqualOp
	GreaterOrEqualOp
	GreaterThanOp
	LessOrEqualOp
	LessThanOp
	EqualTotalOrderOp
	NotEqualTotalOrderOp
	GreaterOrEqualTotalOrderOp
	GreaterThanTotalOrderOp
	LessOrEqualTotalOrderOp
	LessThanTotalOrderOp
	RngBitGeneratorOp
	RngNormalOp
	RngUniformOp
)
