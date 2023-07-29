# Go WASM GPU

This project aims to expose the WebGPU API as Go API that can be used
in wasm projects.

> **Warning:** The project is in early development and the API is likely to
> change in backward incompatible ways!

## Getting Started

You need to add the project as a dependency.

```
go get github.com/mokiat/wasmgpu@latest
```

The implementation uses `syscall/js` calls and as such requires that client
applications are compiled with the `GOOS=js` and `GOARCH=wasm` options.

If you are unfamiliar with how Go and WASM works, then you should have a look at
the official [WebAssembly with Go documentation](https://github.com/golang/go/wiki/WebAssembly).

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
