# Go WASM GPU

This project aims to expose the WebGPU API as Go API that can be used
in wasm projects.

> **Warning:** The project is in early development and the API is likely to
> change in backward incompatible ways!

## Getting Started

You need to add the project as a dependency.

```sh
go get github.com/mokiat/wasmgpu@latest
```

The implementation uses `syscall/js` calls and as such requires that client
applications are compiled with the `GOOS=js` and `GOARCH=wasm` options.

If you are unfamiliar with how Go and WASM works, then you should have a look at
the official [WebAssembly with Go documentation](https://github.com/golang/go/wiki/WebAssembly).


### Example

Following is an example Go code that clears the canvas with a green color.

```go
package main

import (
	"syscall/js"

	"github.com/mokiat/gog/opt"
	"github.com/mokiat/wasmgpu"
)

func main() {
	// You need to ensure that getContext and getDevice methods exist.
	jsContext := js.Global().Call("getContext")
	jsDevice := js.Global().Call("getDevice")

	context := wasmgpu.NewCanvasContext(jsContext)

	device := wasmgpu.NewDevice(jsDevice)
	commandEncoder := device.CreateCommandEncoder()

	renderPass := commandEncoder.BeginRenderPass(wasmgpu.GPURenderPassDescriptor{
		ColorAttachments: []wasmgpu.GPURenderPassColorAttachment{
			{
				View: context.GetCurrentTexture().CreateView(),
				ClearValue: opt.V(wasmgpu.GPUColor{
					R: 0.0,
					G: 1.0,
					B: 0.0,
					A: 1.0,
				}),
				LoadOp:  wasmgpu.GPULoadOpClear,
				StoreOp: wasmgpu.GPUStoreOPStore,
			},
		},
	})
	renderPass.End()

	device.Queue().Submit([]wasmgpu.GPUCommandBuffer{
		commandEncoder.Finish(),
	})
}
```

**NOTE:** In order to get this working, the code needs access to `getContext` and
`getDevice` functions that don't exist by default. You need to expose those
from your page's JavaScript.

```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Experiment</title>
    <script type="text/javascript" src="wasm_exec.js"></script>
  </head>
  <body>
    <canvas width="512" height="512"></canvas>
    <script type="module">
      // Prepare a Canvas with WebGPU support.
      if (!navigator.gpu) {
        throw new Error("WebGPU not suppored :(");
      }
      const adapter = await navigator.gpu.requestAdapter({
        powerPreference: "high-performance",
      });
      if (!adapter) {
        throw new Error("No WebGPU adapter available :(");
      }
      const device = await adapter.requestDevice();
      if (!device) {
        throw new Error("No Device available :(");
      }
      const canvas = document.querySelector("canvas");
      const context = canvas.getContext("webgpu");
      const canvasFormat = navigator.gpu.getPreferredCanvasFormat();
      context.configure({
        device: device,
        format: canvasFormat,
      });

      // The following functions are needed by the Go code to get access
      // to the context and device.
      window.getDevice = () => {
        return device;
      };
      window.getContext = () => {
        return context;
      };

      // This is the standard code to get WebAssembly going with Go.
      const go = new Go();
      const result = await WebAssembly.instantiateStreaming(
        fetch("main.wasm"),
        go.importObject
      );
      await go.run(result.instance);
    </script>
  </body>
</html>
```

## Extending the API

If something is not available in the API, in some cases, you can extend it as a
workaround by extracting the `js.Value` object using `obj.ToJS().(js.Value)`
and assigning it to a wrapper structure.

**Example:**

```go
func ExtendQueue(queue GPUQueue) GPUQueueExt {
  return GPUQueueExt {
    GPUQueue: queue,
    jsValue:  queue.ToJS().(js.Value),
  }
}

type GPUQueueExt struct {
  GPUQueue
  jsValue js.Value
}

func (g GPUQueueExt) DoSomethingNew() {
  g.jsValue.Call("somethingNew")
}
```

In the meantime you can open a Github Issue so that the actual implementation
can be adjusted.

## Contributing

The project is still very early in the design phase. At this point in time
Issues would be the preferred contribution mechanism.
