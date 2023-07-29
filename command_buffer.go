package wasmgpu

import "syscall/js"

// GPUCommandBuffer as described:
// https://gpuweb.github.io/gpuweb/#gpucommandbuffer
type GPUCommandBuffer struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUCommandBuffer) ToJS() any {
	return g.jsValue
}
