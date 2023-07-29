package wasmgpu

import "syscall/js"

// NewCanvasContext creates a new GPUCanvasContext using the specified
// JavaScript reference as the underlying context.
func NewCanvasContext(jsValue js.Value) GPUCanvasContext {
	return GPUCanvasContext{
		jsValue: jsValue,
	}
}

// GPUCanvasContext as described:
// https://gpuweb.github.io/gpuweb/#gpucanvascontext
type GPUCanvasContext struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUCanvasContext) ToJS() any {
	return g.jsValue
}

// GetCurrentTexture as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucanvascontext-getcurrenttexture
func (g GPUCanvasContext) GetCurrentTexture() GPUTexture {
	jsTexture := g.jsValue.Call("getCurrentTexture")
	return GPUTexture{
		jsValue: jsTexture,
	}
}
