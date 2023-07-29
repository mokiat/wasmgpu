package wasmgpu

import "syscall/js"

// NewDevice creates a new GPUDevice that uses the specified JavaScript
// reference of the device.
func NewDevice(jsValue js.Value) GPUDevice {
	return GPUDevice{
		jsValue: jsValue,
	}
}

// GPUDevice as described:
// https://gpuweb.github.io/gpuweb/#gpudevice
type GPUDevice struct {
	jsValue js.Value
}

// ToJS converts this type to one that can be passed as an argument
// to JavaScript.
func (g GPUDevice) ToJS() any {
	return g.jsValue
}

// Queue as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-queue
func (d GPUDevice) Queue() GPUQueue {
	jsQueue := d.jsValue.Get("queue")
	return GPUQueue{
		jsValue: jsQueue,
	}
}

// CreateCommandEncoder as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createcommandencoder
func (d GPUDevice) CreateCommandEncoder() GPUCommandEncoder {
	jsEncoder := d.jsValue.Call("createCommandEncoder")
	return GPUCommandEncoder{
		jsValue: jsEncoder,
	}
}

// CreateBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbuffer
func (g GPUDevice) CreateBuffer(descriptor GPUBufferDescriptor) GPUBuffer {
	jsBuffer := g.jsValue.Call("createBuffer", descriptor.ToJS())
	return GPUBuffer{
		jsValue: jsBuffer,
	}
}
