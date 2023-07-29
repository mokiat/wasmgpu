package wasmgpu

// GPULoadOp as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpuloadop
type GPULoadOp string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPULoadOp) ToJS() any {
	return string(g)
}

const (
	GPULoadOpLoad  GPULoadOp = "load"
	GPULoadOpClear GPULoadOp = "clear"
)

// GPUStoreOp as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpustoreop
type GPUStoreOp string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUStoreOp) ToJS() any {
	return string(g)
}

const (
	GPUStoreOPStore   GPUStoreOp = "store"
	GPUStoreOpDiscard GPUStoreOp = "discard"
)

// GPUVertexStepMode as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpuvertexstepmode
type GPUVertexStepMode string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUVertexStepMode) ToJS() any {
	return string(g)
}

const (
	GPUVertexStepModeVertex   GPUVertexStepMode = "vertex"
	GPUVertexStepModeInstance GPUVertexStepMode = "instance"
)

// GPUVertexFormat as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpuvertexformat
type GPUVertexFormat string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUVertexFormat) ToJS() any {
	return string(g)
}

const (
	GPUVertexFormatUint8x2   GPUVertexFormat = "uint8x2"
	GPUVertexFormatUint8x4   GPUVertexFormat = "uint8x4"
	GPUVertexFormatSint8x2   GPUVertexFormat = "sint8x2"
	GPUVertexFormatSint8x4   GPUVertexFormat = "sint8x4"
	GPUVertexFormatUnorm8x2  GPUVertexFormat = "unorm8x2"
	GPUVertexFormatUnorm8x4  GPUVertexFormat = "unorm8x4"
	GPUVertexFormatSnorm8x2  GPUVertexFormat = "snorm8x2"
	GPUVertexFormatSnorm8x4  GPUVertexFormat = "snorm8x4"
	GPUVertexFormatUint16x2  GPUVertexFormat = "uint16x2"
	GPUVertexFormatUint16x4  GPUVertexFormat = "uint16x4"
	GPUVertexFormatSint16x2  GPUVertexFormat = "sint16x2"
	GPUVertexFormatSint16x4  GPUVertexFormat = "sint16x4"
	GPUVertexFormatUnorm16x2 GPUVertexFormat = "unorm16x2"
	GPUVertexFormatUnorm16x4 GPUVertexFormat = "unorm16x4"
	GPUVertexFormatSnorm16x2 GPUVertexFormat = "snorm16x2"
	GPUVertexFormatSnorm16x4 GPUVertexFormat = "snorm16x4"
	GPUVertexFormatFloat16x2 GPUVertexFormat = "float16x2"
	GPUVertexFormatFloat16x4 GPUVertexFormat = "float16x4"
	GPUVertexFormatFloat32   GPUVertexFormat = "float32"
	GPUVertexFormatFloat32x2 GPUVertexFormat = "float32x2"
	GPUVertexFormatFloat32x3 GPUVertexFormat = "float32x3"
	GPUVertexFormatFloat32x4 GPUVertexFormat = "float32x4"
	GPUVertexFormatUint32    GPUVertexFormat = "uint32"
	GPUVertexFormatUint32x2  GPUVertexFormat = "uint32x2"
	GPUVertexFormatUint32x3  GPUVertexFormat = "uint32x3"
	GPUVertexFormatUint32x4  GPUVertexFormat = "uint32x4"
	GPUVertexFormatSint32    GPUVertexFormat = "sint32"
	GPUVertexFormatSint32x2  GPUVertexFormat = "sint32x2"
	GPUVertexFormatSint32x3  GPUVertexFormat = "sint32x3"
	GPUVertexFormatSint32x4  GPUVertexFormat = "sint32x4"
)
