package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
	"github.com/mokiat/gog/opt"
)

// GPUBufferBinding as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubufferbinding
type GPUBufferBinding struct {
	Buffer GPUBuffer
	Offset opt.T[GPUSize64]
	Size   opt.T[GPUSize64]
}

var _ GPUBindingResource = GPUBufferBinding{}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBufferBinding) ToJS() any {
	result := make(map[string]any)
	result["buffer"] = g.Buffer.ToJS()
	if g.Offset.Specified {
		result["offset"] = g.Offset.Value.ToJS()
	}
	if g.Size.Specified {
		result["size"] = g.Size.Value.ToJS()
	}
	return result
}

func (g GPUBufferBinding) _isGPUBindingResource() {}

// GPUBindGroupLayout as described:
// https://gpuweb.github.io/gpuweb/#gpubindgrouplayout
type GPUBindGroupLayout struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBindGroupLayout) ToJS() any {
	return g.jsValue
}

// GPUBindingResource as described:
// https://gpuweb.github.io/gpuweb/#typedefdef-gpubindingresource
type GPUBindingResource interface {
	_isGPUBindingResource()
	ToJS() any
}

// GPUBindGroupEntry as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgroupentry
type GPUBindGroupEntry struct {
	Binding  GPUIndex32
	Resource GPUBindingResource
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBindGroupEntry) ToJS() any {
	return map[string]any{
		"binding":  g.Binding.ToJS(),
		"resource": g.Resource.ToJS(),
	}
}

// GPUBindGroupDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgroupdescriptor
type GPUBindGroupDescriptor struct {
	Layout  GPUBindGroupLayout
	Entries []GPUBindGroupEntry
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBindGroupDescriptor) ToJS() any {
	return map[string]any{
		"layout": g.Layout.ToJS(),
		"entries": gog.Map(g.Entries, func(entry GPUBindGroupEntry) any {
			return entry.ToJS()
		}),
	}
}

// GPUBindGroup as described:
// https://gpuweb.github.io/gpuweb/#gpubindgroup
type GPUBindGroup struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBindGroup) ToJS() any {
	return g.jsValue
}
