package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog/opt"
)

// GPUCommandEncoder as described:
// https://gpuweb.github.io/gpuweb/#gpucommandencoder
type GPUCommandEncoder struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUCommandEncoder) ToJS() any {
	return g.jsValue
}

// BeginRenderPass as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-beginrenderpass
func (g GPUCommandEncoder) BeginRenderPass(descriptor GPURenderPassDescriptor) GPURenderPassEncoder {
	jsRenderPass := g.jsValue.Call("beginRenderPass", descriptor.ToJS())
	return GPURenderPassEncoder{
		jsValue: jsRenderPass,
	}
}

// BeginComputePass as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-begincomputepass
func (g GPUCommandEncoder) BeginComputePass(descriptor opt.T[GPUComputePassDescriptor]) GPUComputePassEncoder {
	params := make([]any, 1)
	if descriptor.Specified {
		params[0] = descriptor.Value.ToJS()
	} else {
		params[0] = js.Undefined()
	}
	jsComputePass := g.jsValue.Call("beginComputePass", params...)
	return GPUComputePassEncoder{
		jsValue: jsComputePass,
	}
}

// Finish as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucommandencoder-finish
func (g GPUCommandEncoder) Finish() GPUCommandBuffer {
	jsBuffer := g.jsValue.Call("finish")
	return GPUCommandBuffer{
		jsValue: jsBuffer,
	}
}
