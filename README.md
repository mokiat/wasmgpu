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
