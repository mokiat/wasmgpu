package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
)

// GPUComputePassDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpucomputepassdescriptor
type GPUComputePassDescriptor struct{}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUComputePassDescriptor) ToJS() any {
	return map[string]any{}
}

// GPUComputePassEncoder as described:
// https://gpuweb.github.io/gpuweb/#gpucomputepassencoder
type GPUComputePassEncoder struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUComputePassEncoder) ToJS() any {
	return g.jsValue
}

// SetPipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucomputepassencoder-setpipeline
func (g GPUComputePassEncoder) SetPipeline(pipeline GPUComputePipeline) {
	g.jsValue.Call("setPipeline", pipeline.ToJS())
}

// SetBindGroup as described:
// https://gpuweb.github.io/gpuweb/#dom-gpubindingcommandsmixin-setbindgroup
func (g GPUComputePassEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets []GPUBufferDynamicOffset) {
	params := make([]any, 3)
	params[0] = index.ToJS()
	params[1] = bindGroup.ToJS()
	params[2] = gog.Map(dynamicOffsets, func(offset GPUBufferDynamicOffset) any {
		return offset.ToJS()
	})
	g.jsValue.Call("setBindGroup", params...)
}

// DispatchWorkgroups as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucomputepassencoder-dispatchworkgroups
func (g GPUComputePassEncoder) DispatchWorkgroups(workgroupCountX, workgroupCountY, workgroupCountZ GPUSize32) {
	params := make([]any, 3)
	params[0] = workgroupCountX.ToJS()
	if workgroupCountY > 0 {
		params[1] = workgroupCountY.ToJS()
	} else {
		params[1] = js.Undefined()
	}
	if workgroupCountZ > 0 {
		params[2] = workgroupCountZ.ToJS()
	} else {
		params[2] = js.Undefined()
	}
	g.jsValue.Call("dispatchWorkgroups", params...)
}

// End as described:
// https://gpuweb.github.io/gpuweb/#dom-gpucomputepassencoder-end
func (g GPUComputePassEncoder) End() {
	g.jsValue.Call("end")
}
