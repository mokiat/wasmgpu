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

// End as described:
// https://gpuweb.github.io/gpuweb/#dom-gpurenderpassencoder-end
func (g GPURenderPassEncoder) End() {
	g.jsValue.Call("end")
}
