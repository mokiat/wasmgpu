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
