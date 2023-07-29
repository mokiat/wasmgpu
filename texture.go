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

// GPUTexture as described:
// https://gpuweb.github.io/gpuweb/#gputexture
type GPUTexture struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUTexture) ToJS() any {
	return g.jsValue
}

// Format as described:
// https://gpuweb.github.io/gpuweb/#dom-gputexture-format
func (g GPUTexture) Format() GPUTextureFormat {
	jsFormat := g.jsValue.Get("format")
	return GPUTextureFormat(jsFormat.String())
}

// CreateView as described:
// https://gpuweb.github.io/gpuweb/#dom-gputexture-createview
func (g GPUTexture) CreateView() GPUTextureView {
	jsView := g.jsValue.Call("createView")
	return GPUTextureView{
		jsValue: jsView,
	}
}
