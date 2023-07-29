package wasmgpu

import "syscall/js"

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
