package wasmgpu

import "syscall/js"

// GPUBufferDescriptor as described:
// https://gpuweb.github.io/gpuweb/#gpubufferdescriptor
type GPUBufferDescriptor struct {
	Size  GPUSize64
	Usage GPUBufferUsageFlags
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBufferDescriptor) ToJS() any {
	return map[string]any{
		"size":  g.Size.ToJS(),
		"usage": g.Usage.ToJS(),
	}
}

// GPUBuffer as described:
// https://gpuweb.github.io/gpuweb/#gpubuffer
type GPUBuffer struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBuffer) ToJS() any {
	return g.jsValue
}

// Destroy as described:
// https://gpuweb.github.io/gpuweb/#dom-gpubuffer-destroy
func (g GPUBuffer) Destroy() {
	g.jsValue.Call("destroy")
}
