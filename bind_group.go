package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
	"github.com/mokiat/gog/opt"
)

// GPUBufferBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubufferbindinglayout
type GPUBufferBindingLayout struct {
	Type             opt.T[GPUBufferBindingType]
	HasDynamicOffset opt.T[bool]
	MinBindingSize   opt.T[GPUSize64]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBufferBindingLayout) ToJS() any {
	result := make(map[string]any)
	if g.Type.Specified {
		result["type"] = g.Type.Value.ToJS()
	}
	if g.HasDynamicOffset.Specified {
		result["hasDynamicOffset"] = g.HasDynamicOffset.Value
	}
	if g.MinBindingSize.Specified {
		result["minBindingSize"] = g.MinBindingSize.Value.ToJS()
	}
	return result
}

// GPUSamplerBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpusamplerbindinglayout
type GPUSamplerBindingLayout struct {
	Type opt.T[GPUSamplerBindingType]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUSamplerBindingLayout) ToJS() any {
	result := make(map[string]any)
	if g.Type.Specified {
		result["type"] = g.Type.Value.ToJS()
	}
	return result
}

// GPUTextureBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gputexturebindinglayout
type GPUTextureBindingLayout struct {
	SampleType    opt.T[GPUTextureSampleType]
	ViewDimension opt.T[GPUTextureViewDimension]
	Multisampled  opt.T[bool]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUTextureBindingLayout) ToJS() any {
	result := make(map[string]any)
	if g.SampleType.Specified {
		result["sampleType"] = g.SampleType.Value.ToJS()
	}
	if g.ViewDimension.Specified {
		result["viewDimension"] = g.ViewDimension.Value.ToJS()
	}
	if g.Multisampled.Specified {
		result["multisampled"] = g.Multisampled.Value
	}
	return result
}

// GPUStorageTextureBindingLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpustoragetexturebindinglayout
type GPUStorageTextureBindingLayout struct {
	Access        opt.T[GPUStorageTextureAccess]
	Format        GPUTextureFormat
	ViewDimension opt.T[GPUTextureViewDimension]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUStorageTextureBindingLayout) ToJS() any {
	result := make(map[string]any)
	if g.Access.Specified {
		result["access"] = g.Access.Value.ToJS()
	}
	result["format"] = g.Format.ToJS()
	if g.ViewDimension.Specified {
		result["viewDimension"] = g.ViewDimension.Value.ToJS()
	}
	return result
}

// GPUExternalTextureBindingLayout as described:
type GPUExternalTextureBindingLayout struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUExternalTextureBindingLayout) ToJS() any {
	return g.jsValue
}

// GPUBindGroupLayoutEntry as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgrouplayoutentry
type GPUBindGroupLayoutEntry struct {
	Binding         GPUIndex32
	Visibility      GPUShaderStageFlags
	Buffer          opt.T[GPUBufferBindingLayout]
	Sampler         opt.T[GPUSamplerBindingLayout]
	Texture         opt.T[GPUTextureBindingLayout]
	StorageTexture  opt.T[GPUStorageTextureBindingLayout]
	ExternalTexture opt.T[GPUExternalTextureBindingLayout]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBindGroupLayoutEntry) ToJS() any {
	result := make(map[string]any)
	result["binding"] = g.Binding.ToJS()
	result["visibility"] = g.Visibility.ToJS()
	if g.Buffer.Specified {
		result["buffer"] = g.Buffer.Value.ToJS()
	}
	if g.Sampler.Specified {
		result["sampler"] = g.Sampler.Value.ToJS()
	}
	if g.Texture.Specified {
		result["texture"] = g.Texture.Value.ToJS()
	}
	if g.StorageTexture.Specified {
		result["storageTexture"] = g.StorageTexture.Value.ToJS()
	}
	if g.ExternalTexture.Specified {
		result["externalTexture"] = g.ExternalTexture.Value.ToJS()
	}
	return result
}

// GPUBindGroupLayoutDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpubindgrouplayoutdescriptor
type GPUBindGroupLayoutDescriptor struct {
	Entries []GPUBindGroupLayoutEntry
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBindGroupLayoutDescriptor) ToJS() any {
	return map[string]any{
		"entries": gog.Map(g.Entries, func(entry GPUBindGroupLayoutEntry) any {
			return entry.ToJS()
		}),
	}
}

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
