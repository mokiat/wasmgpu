package wasmgpu

import (
	"syscall/js"

	"github.com/mokiat/gog"
	"github.com/mokiat/gog/opt"
)

// GPUPipelineLayout as described:
// https://gpuweb.github.io/gpuweb/#gpupipelinelayout
type GPUPipelineLayout struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUPipelineLayout) ToJS() any {
	return g.jsValue
}

// GPUVertexAttribute as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuvertexattribute
type GPUVertexAttribute struct {
	Format         GPUVertexFormat
	Offset         GPUSize64
	ShaderLocation GPUIndex32
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUVertexAttribute) ToJS() any {
	return map[string]any{
		"format":         g.Format.ToJS(),
		"offset":         g.Offset.ToJS(),
		"shaderLocation": g.ShaderLocation.ToJS(),
	}
}

// GPUVertexBufferLayout as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuvertexbufferlayout
type GPUVertexBufferLayout struct {
	ArrayStride GPUSize64
	StepMode    opt.T[GPUVertexStepMode]
	Attributes  []GPUVertexAttribute
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUVertexBufferLayout) ToJS() any {
	result := make(map[string]any)
	result["arrayStride"] = g.ArrayStride.ToJS()
	if g.StepMode.Specified {
		result["stepMode"] = g.StepMode.Value.ToJS()
	}
	result["attributes"] = gog.Map(g.Attributes, func(attrib GPUVertexAttribute) any {
		return attrib.ToJS()
	})
	return result
}

// GPUVertexState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuvertexstate
type GPUVertexState struct {
	Module     GPUShaderModule
	EntryPoint string
	Buffers    []GPUVertexBufferLayout
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUVertexState) ToJS() any {
	return map[string]any{
		"module":     g.Module.ToJS(),
		"entryPoint": g.EntryPoint,
		"buffers": gog.Map(g.Buffers, func(layout GPUVertexBufferLayout) any {
			return layout.ToJS()
		}),
	}
}

// GPUPrimitiveState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpuprimitivestate
type GPUPrimitiveState struct {
	Topology         opt.T[GPUPrimitiveTopology]
	StripIndexFormat opt.T[GPUIndexFormat]
	FrontFace        opt.T[GPUFrontFace]
	CullMode         opt.T[GPUCullMode]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUPrimitiveState) ToJS() any {
	result := make(map[string]any)
	if g.Topology.Specified {
		result["topology"] = g.Topology.Value.ToJS()
	}
	if g.StripIndexFormat.Specified {
		result["stripIndexFormat"] = g.StripIndexFormat.Value.ToJS()
	}
	if g.FrontFace.Specified {
		result["frontFace"] = g.FrontFace.Value.ToJS()
	}
	if g.CullMode.Specified {
		result["cullMode"] = g.CullMode.Value.ToJS()
	}
	return result
}

// GPUStencilFaceState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpustencilfacestate
type GPUStencilFaceState struct {
	Compare     opt.T[GPUCompareFunction]
	FailOp      opt.T[GPUStencilOperation]
	DepthFailOp opt.T[GPUStencilOperation]
	PassOp      opt.T[GPUStencilOperation]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUStencilFaceState) ToJS() any {
	result := make(map[string]any)
	if g.Compare.Specified {
		result["compare"] = g.Compare.Value.ToJS()
	}
	if g.FailOp.Specified {
		result["failOp"] = g.FailOp.Value.ToJS()
	}
	if g.DepthFailOp.Specified {
		result["depthFailOp"] = g.DepthFailOp.Value.ToJS()
	}
	if g.PassOp.Specified {
		result["passOp"] = g.PassOp.Value.ToJS()
	}
	return result
}

// GPUDepthStencilState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpudepthstencilstate
type GPUDepthStencilState struct {
	Format              GPUTextureFormat
	DepthWriteEnabled   bool
	DepthCompare        GPUCompareFunction
	StencilFront        opt.T[GPUStencilFaceState]
	StencilBack         opt.T[GPUStencilFaceState]
	StencilReadMask     opt.T[GPUStencilValue]
	StencilWriteMask    opt.T[GPUStencilValue]
	DepthBias           opt.T[GPUDepthBias]
	DepthBiasSlopeScale opt.T[float32]
	DepthBiasClamp      opt.T[float32]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUDepthStencilState) ToJS() any {
	result := make(map[string]any)
	result["format"] = g.Format.ToJS()
	result["depthWriteEnabled"] = g.DepthWriteEnabled
	result["depthCompare"] = g.DepthCompare.ToJS()
	if g.StencilFront.Specified {
		result["stencilFront"] = g.StencilFront.Value.ToJS()
	}
	if g.StencilBack.Specified {
		result["stencilBack"] = g.StencilBack.Value.ToJS()
	}
	if g.StencilReadMask.Specified {
		result["stencilReadMask"] = g.StencilReadMask.Value.ToJS()
	}
	if g.StencilWriteMask.Specified {
		result["stencilWriteMask"] = g.StencilWriteMask.Value.ToJS()
	}
	if g.DepthBias.Specified {
		result["depthBias"] = g.DepthBias.Value.ToJS()
	}
	if g.DepthBiasSlopeScale.Specified {
		result["depthBiasSlopeScale"] = g.DepthBiasSlopeScale.Value
	}
	if g.DepthBiasClamp.Specified {
		result["depthBiasClamp"] = g.DepthBiasClamp.Value
	}
	return result
}

// GPUMultisampleState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpumultisamplestate
type GPUMultisampleState struct {
	Count                  opt.T[GPUSize32]
	Mask                   opt.T[GPUSampleMask]
	AlphaToCoverageEnabled opt.T[bool]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUMultisampleState) ToJS() any {
	result := make(map[string]any)
	if g.Count.Specified {
		result["count"] = g.Count.Value.ToJS()
	}
	if g.Mask.Specified {
		result["mask"] = g.Mask.Value.ToJS()
	}
	if g.AlphaToCoverageEnabled.Specified {
		result["alphaToCoverageEnabled"] = g.AlphaToCoverageEnabled.Value
	}
	return result
}

// GPUBlendComponent as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpublendcomponent
type GPUBlendComponent struct {
	Operation opt.T[GPUBlendOperation]
	SrcFactor opt.T[GPUBlendFactor]
	DstFactor opt.T[GPUBlendFactor]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBlendComponent) ToJS() any {
	result := make(map[string]any)
	if g.Operation.Specified {
		result["operation"] = g.Operation.Value.ToJS()
	}
	if g.SrcFactor.Specified {
		result["srcFactor"] = g.SrcFactor.Value.ToJS()
	}
	if g.DstFactor.Specified {
		result["dstFactor"] = g.DstFactor.Value.ToJS()
	}
	return result
}

// GPUBlendState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpublendstate
type GPUBlendState struct {
	Color GPUBlendComponent
	Alpha GPUBlendComponent
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUBlendState) ToJS() any {
	return map[string]any{
		"color": g.Color.ToJS(),
		"alpha": g.Alpha.ToJS(),
	}
}

// GPUColorTargetState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpucolortargetstate
type GPUColorTargetState struct {
	Format    GPUTextureFormat
	Blend     opt.T[GPUBlendState]
	WriteMask opt.T[GPUColorWriteFlags]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUColorTargetState) ToJS() any {
	result := make(map[string]any)
	result["format"] = g.Format.ToJS()
	if g.Blend.Specified {
		result["blend"] = g.Blend.Value.ToJS()
	}
	if g.WriteMask.Specified {
		result["writeMask"] = g.WriteMask.Value.ToJS()
	}
	return result
}

// GPUFragmentState as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpufragmentstate
type GPUFragmentState struct {
	Module     GPUShaderModule
	EntryPoint string
	Targets    []GPUColorTargetState
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUFragmentState) ToJS() any {
	return map[string]any{
		"module":     g.Module.ToJS(),
		"entryPoint": g.EntryPoint,
		"targets": gog.Map(g.Targets, func(target GPUColorTargetState) any {
			return target.ToJS()
		}),
	}
}

// GPURenderPipelineDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpurenderpipelinedescriptor
type GPURenderPipelineDescriptor struct {
	Layout       opt.T[GPUPipelineLayout]
	Vertex       GPUVertexState
	Primitive    opt.T[GPUPrimitiveState]
	DepthStencil opt.T[GPUDepthStencilState]
	Multisample  opt.T[GPUMultisampleState]
	Fragment     opt.T[GPUFragmentState]
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPURenderPipelineDescriptor) ToJS() any {
	result := make(map[string]any)
	if g.Layout.Specified {
		result["layout"] = g.Layout.Value.ToJS()
	} else {
		result["layout"] = "auto"
	}
	result["vertex"] = g.Vertex.ToJS()
	if g.Primitive.Specified {
		result["primitive"] = g.Primitive.Value.ToJS()
	}
	if g.DepthStencil.Specified {
		result["depthStencil"] = g.DepthStencil.Value.ToJS()
	}
	if g.Multisample.Specified {
		result["multisample"] = g.Multisample.Value.ToJS()
	}
	if g.Fragment.Specified {
		result["fragment"] = g.Fragment.Value.ToJS()
	}
	return result
}

// GPURenderPipeline as described:
// https://gpuweb.github.io/gpuweb/#gpurenderpipeline
type GPURenderPipeline struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPURenderPipeline) ToJS() any {
	return g.jsValue
}
