package wasmgpu

import "syscall/js"

// GPUProgrammableStage as described:
// https://gpuweb.github.io/gpuweb/#gpuprogrammablestage
type GPUProgrammableStage struct {
	Module     GPUShaderModule
	EntryPoint string
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUProgrammableStage) ToJS() any {
	return map[string]any{
		"module":     g.Module.ToJS(),
		"entryPoint": g.EntryPoint,
	}
}

// GPUShaderModuleDescriptor as described:
// https://gpuweb.github.io/gpuweb/#dictdef-gpushadermoduledescriptor
type GPUShaderModuleDescriptor struct {
	Code string
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUShaderModuleDescriptor) ToJS() any {
	return map[string]any{
		"code": g.Code,
	}
}

// GPUShaderModule as described:
// https://gpuweb.github.io/gpuweb/#gpushadermodule
type GPUShaderModule struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUShaderModule) ToJS() any {
	return g.jsValue
}
