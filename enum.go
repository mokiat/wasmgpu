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

// GPUBufferUsageFlags as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpubufferusageflags
type GPUBufferUsageFlags GPUFlagsConstant

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBufferUsageFlags) ToJS() any {
	return uint32(g)
}

const (
	GPUBufferUsageFlagsMapRead      GPUBufferUsageFlags = 0x0001
	GPUBufferUsageFlagsMapWrite     GPUBufferUsageFlags = 0x0002
	GPUBufferUsageFlagsCopySrc      GPUBufferUsageFlags = 0x0004
	GPUBufferUsageFlagsCopyDst      GPUBufferUsageFlags = 0x0008
	GPUBufferUsageFlagsIndex        GPUBufferUsageFlags = 0x0010
	GPUBufferUsageFlagsVertex       GPUBufferUsageFlags = 0x0020
	GPUBufferUsageFlagsUniform      GPUBufferUsageFlags = 0x0040
	GPUBufferUsageFlagsStorage      GPUBufferUsageFlags = 0x0080
	GPUBufferUsageFlagsIndirect     GPUBufferUsageFlags = 0x0100
	GPUBufferUsageFlagsQueryResolve GPUBufferUsageFlags = 0x0200
)

// GPUTextureFormat as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gputextureformat
type GPUTextureFormat string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUTextureFormat) ToJS() any {
	return string(g)
}

const (
	GPUTextureFormatR8Unorm GPUTextureFormat = "r8unorm"
	GPUTextureFormatR8Snorm GPUTextureFormat = "r8snorm"
	GPUTextureFormatR8Uint  GPUTextureFormat = "r8uint"
	GPUTextureFormatR8Sint  GPUTextureFormat = "r8sint"

	GPUTextureFormatR16Uint  GPUTextureFormat = "r16uint"
	GPUTextureFormatR16Sint  GPUTextureFormat = "r16sint"
	GPUTextureFormatR16float GPUTextureFormat = "r16float"
	GPUTextureFormatRG8Unorm GPUTextureFormat = "rg8unorm"
	GPUTextureFormatRG8Snorm GPUTextureFormat = "rg8snorm"
	GPUTextureFormatRG8Uint  GPUTextureFormat = "rg8uint"
	GPUTextureFormatRG8Sint  GPUTextureFormat = "rg8sint"

	GPUTextureFormatR32Uint        GPUTextureFormat = "r32uint"
	GPUTextureFormatR32Sint        GPUTextureFormat = "r32sint"
	GPUTextureFormatR32float       GPUTextureFormat = "r32float"
	GPUTextureFormatRG16Uint       GPUTextureFormat = "rg16uint"
	GPUTextureFormatRG16Sint       GPUTextureFormat = "rg16sint"
	GPUTextureFormatRG16float      GPUTextureFormat = "rg16float"
	GPUTextureFormatRGBA8Unorm     GPUTextureFormat = "rgba8unorm"
	GPUTextureFormatRGBA8UnormSRGB GPUTextureFormat = "rgba8unorm-srgb"
	GPUTextureFormatRGBA8Snorm     GPUTextureFormat = "rgba8snorm"
	GPUTextureFormatRGBA8Uint      GPUTextureFormat = "rgba8uint"
	GPUTextureFormatRGBA8Sint      GPUTextureFormat = "rgba8sint"
	GPUTextureFormatBGRA8Unorm     GPUTextureFormat = "bgra8unorm"
	GPUTextureFormatBGRA8UnormSRGB GPUTextureFormat = "bgra8unorm-srgb"
	GPUTextureFormatRGB9E5UFloat   GPUTextureFormat = "rgb9e5ufloat"
	GPUTextureFormatRGB10A2Uint    GPUTextureFormat = "rgb10a2uint"
	GPUTextureFormatRGB10A2Unorm   GPUTextureFormat = "rgb10a2unorm"
	GPUTextureFormatRG11B10Ufloat  GPUTextureFormat = "rg11b10ufloat"

	GPUTextureFormatRG32Uint    GPUTextureFormat = "rg32uint"
	GPUTextureFormatRG32Sint    GPUTextureFormat = "rg32sint"
	GPUTextureFormatRG32Float   GPUTextureFormat = "rg32float"
	GPUTextureFormatRGBA16Uint  GPUTextureFormat = "rgba16uint"
	GPUTextureFormatRGBA16Sint  GPUTextureFormat = "rgba16sint"
	GPUTextureFormatRGBA16Float GPUTextureFormat = "rgba16float"

	GPUTextureFormatRGBA32Uint  GPUTextureFormat = "rgba32uint"
	GPUTextureFormatRGBA32Sint  GPUTextureFormat = "rgba32sint"
	GPUTextureFormatRGBA32Float GPUTextureFormat = "rgba32float"

	GPUTextureFormatStencil8            GPUTextureFormat = "stencil8"
	GPUTextureFormatDepth16Unorm        GPUTextureFormat = "depth16unorm"
	GPUTextureFormatDepth24Plus         GPUTextureFormat = "depth24plus"
	GPUTextureFormatDepth24PlusStencil8 GPUTextureFormat = "depth24plus-stencil8"
	GPUTextureFormatDepth32Float        GPUTextureFormat = "depth32float"

	GPUTextureFormatDepth32FloatStencil8 GPUTextureFormat = "depth32float-stencil8"
)

// GPUBlendOperation as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpublendoperation
type GPUBlendOperation string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBlendOperation) ToJS() any {
	return string(g)
}

const (
	GPUBlendOperationAdd             GPUBlendOperation = "add"
	GPUBlendOperationSubtract        GPUBlendOperation = "subtract"
	GPUBlendOperationReverseSubtract GPUBlendOperation = "reverse-subtract"
	GPUBlendOperationMin             GPUBlendOperation = "min"
	GPUBlendOperationMax             GPUBlendOperation = "max"
)

// GPUBlendFactor as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpublendfactor
type GPUBlendFactor string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBlendFactor) ToJS() any {
	return string(g)
}

const (
	GPUBlendFactorZero              GPUBlendFactor = "zero"
	GPUBlendFactorOne               GPUBlendFactor = "one"
	GPUBlendFactorSrc               GPUBlendFactor = "src"
	GPUBlendFactorOneMinusSrc       GPUBlendFactor = "one-minus-src"
	GPUBlendFactorSrcAlpha          GPUBlendFactor = "src-alpha"
	GPUBlendFactorOneMinusSrcAlpha  GPUBlendFactor = "one-minus-src-alpha"
	GPUBlendFactorDst               GPUBlendFactor = "dst"
	GPUBlendFactorOneMinusDst       GPUBlendFactor = "one-minus-dst"
	GPUBlendFactorDstAlpha          GPUBlendFactor = "dst-alpha"
	GPUBlendFactorOneMinusDstAlpha  GPUBlendFactor = "one-minus-dst-alpha"
	GPUBlendFactorSrcAlphaSaturated GPUBlendFactor = "src-alpha-saturated"
	GPUBlendFactorConstant          GPUBlendFactor = "constant"
	GPUBlendFactorOneMinusConstant  GPUBlendFactor = "one-minus-constant"
)

// GPUColorWriteFlags as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpucolorwriteflags
type GPUColorWriteFlags GPUFlagsConstant

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUColorWriteFlags) ToJS() any {
	return uint32(g)
}

const (
	GPUColorWriteFlagsRed   GPUColorWriteFlags = 0x1
	GPUColorWriteFlagsGreen GPUColorWriteFlags = 0x2
	GPUColorWriteFlagsBlue  GPUColorWriteFlags = 0x4
	GPUColorWriteFlagsAlpha GPUColorWriteFlags = 0x8
	GPUColorWriteFlagsAll   GPUColorWriteFlags = 0xF
)

// GPUPrimitiveTopology as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpuprimitivetopology
type GPUPrimitiveTopology string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUPrimitiveTopology) ToJS() any {
	return string(g)
}

const (
	GPUPrimitiveTopologyPointList     GPUPrimitiveTopology = "point-list"
	GPUPrimitiveTopologyLineList      GPUPrimitiveTopology = "line-list"
	GPUPrimitiveTopologyLineStrip     GPUPrimitiveTopology = "line-strip"
	GPUPrimitiveTopologyTriangleList  GPUPrimitiveTopology = "triangle-list"
	GPUPrimitiveTopologyTriangleStrip GPUPrimitiveTopology = "triangle-strip"
)

// GPUIndexFormat as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpuindexformat
type GPUIndexFormat string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUIndexFormat) ToJS() any {
	return string(g)
}

const (
	GPUIndexFormatUint16 GPUIndexFormat = "uint16"
	GPUIndexFormatUint32 GPUIndexFormat = "uint32"
)

// GPUFrontFace as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpufrontface
type GPUFrontFace string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUFrontFace) ToJS() any {
	return string(g)
}

const (
	GPUFrontFaceCCW GPUFrontFace = "ccw"
	GPUFrontFaceCW  GPUFrontFace = "cw"
)

// GPUCullMode as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpucullmode
type GPUCullMode string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUCullMode) ToJS() any {
	return string(g)
}

const (
	GPUCullModeNone  GPUCullMode = "none"
	GPUCullModeFront GPUCullMode = "front"
	GPUCullModeBack  GPUCullMode = "back"
)

// GPUCompareFunction as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpucomparefunction
type GPUCompareFunction string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUCompareFunction) ToJS() any {
	return string(g)
}

const (
	GPUCompareFunctionNever        GPUCompareFunction = "never"
	GPUCompareFunctionLess         GPUCompareFunction = "less"
	GPUCompareFunctionEqual        GPUCompareFunction = "equal"
	GPUCompareFunctionLessEqual    GPUCompareFunction = "less-equal"
	GPUCompareFunctionGreater      GPUCompareFunction = "greater"
	GPUCompareFunctionNotEqual     GPUCompareFunction = "not-equal"
	GPUCompareFunctionGreaterEqual GPUCompareFunction = "greater-equal"
	GPUCompareFunctionAlways       GPUCompareFunction = "always"
)

// GPUStencilOperation as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpustenciloperation
type GPUStencilOperation string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUStencilOperation) ToJS() any {
	return string(g)
}

const (
	GPUStencilOperationKeep           GPUStencilOperation = "keep"
	GPUStencilOperationZero           GPUStencilOperation = "zero"
	GPUStencilOperationReplace        GPUStencilOperation = "replace"
	GPUStencilOperationInvert         GPUStencilOperation = "invert"
	GPUStencilOperationIncrementClamp GPUStencilOperation = "increment-clamp"
	GPUStencilOperationDecrementClamp GPUStencilOperation = "decrement-clamp"
	GPUStencilOperationIncrementWrap  GPUStencilOperation = "increment-wrap"
	GPUStencilOperationDecrementWrap  GPUStencilOperation = "decrement-wrap"
)

// GPUShaderStageFlags as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpushaderstageflags
type GPUShaderStageFlags GPUFlagsConstant

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUShaderStageFlags) ToJS() any {
	return uint32(g)
}

const (
	GPUShaderStageFlagsVertex   GPUShaderStageFlags = 0x1
	GPUShaderStageFlagsFragment GPUShaderStageFlags = 0x2
	GPUShaderStageFlagsCompute  GPUShaderStageFlags = 0x4
)

// GPUBufferBindingType as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpubufferbindingtype
type GPUBufferBindingType string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBufferBindingType) ToJS() any {
	return string(g)
}

const (
	GPUBufferBindingTypeUniform         GPUBufferBindingType = "uniform"
	GPUBufferBindingTypeStorage         GPUBufferBindingType = "storage"
	GPUBufferBindingTypeReadOnlyStorage GPUBufferBindingType = "read-only-storage"
)

// GPUSamplerBindingType as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpusamplerbindingtype
type GPUSamplerBindingType string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUSamplerBindingType) ToJS() any {
	return string(g)
}

const (
	GPUSamplerBindingTypeFiltering    GPUSamplerBindingType = "filtering"
	GPUSamplerBindingTypeNonFiltering GPUSamplerBindingType = "non-filtering"
	GPUSamplerBindingTypeComparison   GPUSamplerBindingType = "comparison"
)

// GPUTextureSampleType as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gputexturesampletype
type GPUTextureSampleType string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUTextureSampleType) ToJS() any {
	return string(g)
}

const (
	GPUTextureSampleTypeFloat             GPUTextureSampleType = "float"
	GPUTextureSampleTypeUnfilterableFloat GPUTextureSampleType = "unfilterable-float"
	GPUTextureSampleTypeDepth             GPUTextureSampleType = "depth"
	GPUTextureSampleTypeSint              GPUTextureSampleType = "sint"
	GPUTextureSampleTypeUint              GPUTextureSampleType = "uint"
)

// GPUTextureViewDimension as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gputextureviewdimension
type GPUTextureViewDimension string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUTextureViewDimension) ToJS() any {
	return string(g)
}

const (
	GPUTextureViewDimension1D        GPUTextureViewDimension = "1d"
	GPUTextureViewDimension2D        GPUTextureViewDimension = "2d"
	GPUTextureViewDimension2DArray   GPUTextureViewDimension = "2d-array"
	GPUTextureViewDimensionCube      GPUTextureViewDimension = "cube"
	GPUTextureViewDimensionCubeArray GPUTextureViewDimension = "cube-array"
	GPUTextureViewDimension3D        GPUTextureViewDimension = "3d"
)

// GPUStorageTextureAccess as described:
// https://gpuweb.github.io/gpuweb/#enumdef-gpustoragetextureaccess
type GPUStorageTextureAccess string

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUStorageTextureAccess) ToJS() any {
	return string(g)
}

const (
	GPUStorageTextureAccessWriteOnly GPUStorageTextureAccess = "write-only"
)
