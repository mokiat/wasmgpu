package wasmgpu

import "syscall/js"

// GPUTextureView as described:
// https://gpuweb.github.io/gpuweb/#gputextureview
type GPUTextureView struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUTextureView) ToJS() any {
	return g.jsValue
}
