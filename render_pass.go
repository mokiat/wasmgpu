package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
	"github.com/mokiat/gog/opt"
)

// GPURenderPassColorAttachment as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpurenderpasscolorattachment
type GPURenderPassColorAttachment struct {
	View          GPUTextureView
	ResolveTarget opt.T[GPUTextureView]
	ClearValue    opt.T[GPUColor]
	LoadOp        GPULoadOp
	StoreOp       GPUStoreOp
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPURenderPassColorAttachment) ToJS() any {
	result := make(map[string]any)
	result["view"] = g.View.jsValue
	result["loadOp"] = g.LoadOp.ToJS()
	result["storeOp"] = g.StoreOp.ToJS()
	if g.ClearValue.Specified {
		result["clearValue"] = g.ClearValue.Value.ToJS()
	}
	if g.ResolveTarget.Specified {
		result["resolveTarget"] = g.ResolveTarget.Value.ToJS()
	}
	return result
}

// GPURenderPassDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpurenderpassdescriptor
type GPURenderPassDescriptor struct {
	ColorAttachments []GPURenderPassColorAttachment
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPURenderPassDescriptor) ToJS() any {
	result := make(map[string]any)
	result["colorAttachments"] = gog.Map(g.ColorAttachments, func(attachment GPURenderPassColorAttachment) any {
		return attachment.ToJS()
	})
	return result
}

// GPURenderPassEncoder as described:
// https://gpuweb.github.io/gpuweb/#gpurenderpassencoder
type GPURenderPassEncoder struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPURenderPassEncoder) ToJS() any {
	return g.jsValue
}

// SetPipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpurendercommandsmixin-setpipeline
func (g GPURenderPassEncoder) SetPipeline(pipeline GPURenderPipeline) {
	g.jsValue.Call("setPipeline", pipeline.ToJS())
}

// SetVertexBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpurendercommandsmixin-setvertexbuffer
func (g GPURenderPassEncoder) SetVertexBuffer(slot GPUIndex32, vertexBuffer GPUBuffer, offset, size opt.T[GPUSize64]) {
	params := make([]any, 4)
	params[0] = slot.ToJS()
	params[1] = vertexBuffer.ToJS()
	if offset.Specified {
		params[2] = offset.Value.ToJS()
	} else {
		params[2] = js.Undefined()
	}
	if size.Specified {
		params[3] = size.Value.ToJS()
	} else {
		params[3] = js.Undefined()
	}
	g.jsValue.Call("setVertexBuffer", params...)
}

// SetBindGroup as described:
// https://gpuweb.github.io/gpuweb/#gpubindingcommandsmixin-setbindgroup
func (g GPURenderPassEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets []GPUBufferDynamicOffset) {
	params := make([]any, 3)
	params[0] = index.ToJS()
	params[1] = bindGroup.ToJS()
	params[2] = gog.Map(dynamicOffsets, func(offset GPUBufferDynamicOffset) any {
		return offset.ToJS()
	})
	g.jsValue.Call("setBindGroup", params...)
}

// Draw as described:
// https://gpuweb.github.io/gpuweb/#dom-gpurendercommandsmixin-draw
func (g GPURenderPassEncoder) Draw(vertexCount GPUSize32, instanceCount, firstVertex, firstInstance opt.T[GPUSize32]) {
	params := make([]any, 4)
	params[0] = vertexCount.ToJS()
	if instanceCount.Specified {
		params[1] = instanceCount.Value.ToJS()
	} else {
		params[1] = js.Undefined()
	}
	if firstVertex.Specified {
		params[2] = firstVertex.Value.ToJS()
	} else {
		params[2] = js.Undefined()
	}
	if firstInstance.Specified {
		params[3] = firstInstance.Value.ToJS()
	} else {
		params[3] = js.Undefined()
	}
	g.jsValue.Call("draw", params...)
}

// End as described:
// https://gpuweb.github.io/gpuweb/#dom-gpurenderpassencoder-end
func (g GPURenderPassEncoder) End() {
	g.jsValue.Call("end")
}
