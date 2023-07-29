package wasmgpu

// GPUSize64 as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpusize64
type GPUSize64 uint64

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUSize64) ToJS() any {
	return uint64(g)
}

// GPUIntegerCoordinate as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpuintegercoordinate
type GPUIntegerCoordinate uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUIntegerCoordinate) ToJS() any {
	return uint32(g)
}

// GPUIndex32 as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpuindex32
type GPUIndex32 uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUIndex32) ToJS() any {
	return uint32(g)
}

// GPUSize32 as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpusize32
type GPUSize32 uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUSize32) ToJS() any {
	return uint32(g)
}

// GPUColor as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpucolor
type GPUColor struct {
	R float64
	G float64
	B float64
	A float64
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUColor) ToJS() any {
	return []any{g.R, g.G, g.B, g.A}
}

// GPUFlagsConstant as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpuflagsconstant
type GPUFlagsConstant uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUFlagsConstant) ToJS() any {
	return uint32(g)
}

// GPUStencilValue as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpustencilvalue
type GPUStencilValue uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUStencilValue) ToJS() any {
	return uint32(g)
}

// GPUDepthBias as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpudepthbias
type GPUDepthBias int32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUDepthBias) ToJS() any {
	return int32(g)
}

// GPUSampleMask as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpusamplemask
type GPUSampleMask uint32

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUSampleMask) ToJS() any {
	return uint32(g)
}
