package main

import (
	"bytes"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/mokiat/gblob"
	"github.com/mokiat/gog/opt"
	"github.com/mokiat/wasmgpu"
)

const (
	gridSize       = 32
	updateInterval = 100 * time.Millisecond
)

func main() {
	jsContext := js.Global().Call("getContext")
	jsDevice := js.Global().Call("getDevice")

	context := wasmgpu.NewCanvasContext(jsContext)
	canvasFormat := context.GetCurrentTexture().Format()

	device := wasmgpu.NewDevice(jsDevice)

	vertexData := createQuadVertexData()
	vertexBuffer := device.CreateBuffer(wasmgpu.GPUBufferDescriptor{
		Size:  wasmgpu.GPUSize64(len(vertexData)),
		Usage: wasmgpu.GPUBufferUsageFlagsVertex | wasmgpu.GPUBufferUsageFlagsCopyDst,
	})
	device.Queue().WriteBuffer(vertexBuffer, 0, vertexData)

	uniformData := createGridUniformData()
	uniformBuffer := device.CreateBuffer(wasmgpu.GPUBufferDescriptor{
		Size:  wasmgpu.GPUSize64(len(uniformData)),
		Usage: wasmgpu.GPUBufferUsageFlagsUniform | wasmgpu.GPUBufferUsageFlagsCopyDst,
	})
	device.Queue().WriteBuffer(uniformBuffer, 0, uniformData)

	cellStateData := createCellStateData()

	cellStateStorages := []wasmgpu.GPUBuffer{
		device.CreateBuffer(wasmgpu.GPUBufferDescriptor{
			Size:  wasmgpu.GPUSize64(len(cellStateData)),
			Usage: wasmgpu.GPUBufferUsageFlagsStorage | wasmgpu.GPUBufferUsageFlagsCopyDst,
		}),
		device.CreateBuffer(wasmgpu.GPUBufferDescriptor{
			Size:  wasmgpu.GPUSize64(len(cellStateData)),
			Usage: wasmgpu.GPUBufferUsageFlagsStorage | wasmgpu.GPUBufferUsageFlagsCopyDst,
		}),
	}
	device.Queue().WriteBuffer(cellStateStorages[0], 0, cellStateData)

	bindGroupLayout := device.CreateBindGroupLayout(wasmgpu.GPUBindGroupLayoutDescriptor{
		Entries: []wasmgpu.GPUBindGroupLayoutEntry{
			{
				Binding:    wasmgpu.GPUIndex32(0),
				Visibility: wasmgpu.GPUShaderStageFlagsVertex | wasmgpu.GPUShaderStageFlagsFragment | wasmgpu.GPUShaderStageFlagsCompute,
				Buffer: opt.V(wasmgpu.GPUBufferBindingLayout{
					Type: opt.V(wasmgpu.GPUBufferBindingTypeUniform),
				}),
			},
			{
				Binding:    wasmgpu.GPUIndex32(1),
				Visibility: wasmgpu.GPUShaderStageFlagsVertex | wasmgpu.GPUShaderStageFlagsCompute,
				Buffer: opt.V(wasmgpu.GPUBufferBindingLayout{
					Type: opt.V(wasmgpu.GPUBufferBindingTypeReadOnlyStorage),
				}),
			},
			{
				Binding:    wasmgpu.GPUIndex32(2),
				Visibility: wasmgpu.GPUShaderStageFlagsCompute,
				Buffer: opt.V(wasmgpu.GPUBufferBindingLayout{
					Type: opt.V(wasmgpu.GPUBufferBindingTypeStorage),
				}),
			},
		},
	})

	bindGroups := [2]wasmgpu.GPUBindGroup{
		device.CreateBindGroup(wasmgpu.GPUBindGroupDescriptor{
			Layout: bindGroupLayout,
			Entries: []wasmgpu.GPUBindGroupEntry{
				{
					Binding: wasmgpu.GPUIndex32(0),
					Resource: wasmgpu.GPUBufferBinding{
						Buffer: uniformBuffer,
					},
				},
				{
					Binding: wasmgpu.GPUIndex32(1),
					Resource: wasmgpu.GPUBufferBinding{
						Buffer: cellStateStorages[0],
					},
				},
				{
					Binding: wasmgpu.GPUIndex32(2),
					Resource: wasmgpu.GPUBufferBinding{
						Buffer: cellStateStorages[1],
					},
				},
			},
		}),
		device.CreateBindGroup(wasmgpu.GPUBindGroupDescriptor{
			Layout: bindGroupLayout,
			Entries: []wasmgpu.GPUBindGroupEntry{
				{
					Binding: wasmgpu.GPUIndex32(0),
					Resource: wasmgpu.GPUBufferBinding{
						Buffer: uniformBuffer,
					},
				},
				{
					Binding: wasmgpu.GPUIndex32(1),
					Resource: wasmgpu.GPUBufferBinding{
						Buffer: cellStateStorages[1],
					},
				},
				{
					Binding: wasmgpu.GPUIndex32(2),
					Resource: wasmgpu.GPUBufferBinding{
						Buffer: cellStateStorages[0],
					},
				},
			},
		}),
	}

	pipelineLayout := device.CreatePipelineLayout(wasmgpu.GPUPipelineLayoutDescriptor{
		BindGroupLayouts: []wasmgpu.GPUBindGroupLayout{
			bindGroupLayout,
		},
	})

	renderShaderModule := device.CreateShaderModule(wasmgpu.GPUShaderModuleDescriptor{
		Code: renderShaderCode,
	})

	renderPipeline := device.CreateRenderPipeline(wasmgpu.GPURenderPipelineDescriptor{
		Layout: opt.V(pipelineLayout),
		Vertex: wasmgpu.GPUVertexState{
			Module:     renderShaderModule,
			EntryPoint: "vertexMain",
			Buffers: []wasmgpu.GPUVertexBufferLayout{
				{
					ArrayStride: wasmgpu.GPUSize64(2 * 4),
					Attributes: []wasmgpu.GPUVertexAttribute{
						{
							Format:         wasmgpu.GPUVertexFormatFloat32x2,
							Offset:         wasmgpu.GPUSize64(0),
							ShaderLocation: wasmgpu.GPUIndex32(0),
						},
					},
				},
			},
		},
		Fragment: opt.V(wasmgpu.GPUFragmentState{
			Module:     renderShaderModule,
			EntryPoint: "fragmentMain",
			Targets: []wasmgpu.GPUColorTargetState{
				{
					Format: canvasFormat,
				},
			},
		}),
	})

	simulationShaderModule := device.CreateShaderModule(wasmgpu.GPUShaderModuleDescriptor{
		Code: computeShaderCode,
	})

	simulationPipeline := device.CreateComputePipeline(wasmgpu.GPUComputePipelineDescriptor{
		Layout: opt.V(pipelineLayout),
		Compute: wasmgpu.GPUProgrammableStage{
			Module:     simulationShaderModule,
			EntryPoint: "computeMain",
		},
	})

	bindGroupIndex := 0
	updateGrid := js.FuncOf(func(this js.Value, args []js.Value) any {
		// Start recording commands.
		commandEncoder := device.CreateCommandEncoder()

		// Run simulation through compute shader.
		computePass := commandEncoder.BeginComputePass(opt.Unspecified[wasmgpu.GPUComputePassDescriptor]())
		computePass.SetPipeline(simulationPipeline)
		computePass.SetBindGroup(0, bindGroups[bindGroupIndex%2], []wasmgpu.GPUBufferDynamicOffset{})
		workgroupCount := gridSize / 8 /* workgroupSize */
		computePass.DispatchWorkgroups(wasmgpu.GPUSize32(workgroupCount), wasmgpu.GPUSize32(workgroupCount), wasmgpu.GPUSize32(0))
		computePass.End()

		// Swap bind groups (output of compute shader becomes input of render shader).
		bindGroupIndex++

		// Render cell quads to the screen.
		renderPass := commandEncoder.BeginRenderPass(wasmgpu.GPURenderPassDescriptor{
			ColorAttachments: []wasmgpu.GPURenderPassColorAttachment{
				{
					View: context.GetCurrentTexture().CreateView(),
					ClearValue: opt.V(wasmgpu.GPUColor{
						R: 0.0,
						G: 0.1,
						B: 0.3,
						A: 1.0,
					}),
					LoadOp:  wasmgpu.GPULoadOpClear,
					StoreOp: wasmgpu.GPUStoreOPStore,
				},
			},
		})
		renderPass.SetPipeline(renderPipeline)
		renderPass.SetVertexBuffer(0, vertexBuffer, opt.Unspecified[wasmgpu.GPUSize64](), opt.Unspecified[wasmgpu.GPUSize64]())
		renderPass.SetBindGroup(0, bindGroups[bindGroupIndex%2], []wasmgpu.GPUBufferDynamicOffset{})
		renderPass.Draw(6, opt.V[wasmgpu.GPUSize32](gridSize*gridSize), opt.Unspecified[wasmgpu.GPUSize32](), opt.Unspecified[wasmgpu.GPUSize32]())
		renderPass.End()

		// Submit recorded commands.
		device.Queue().Submit([]wasmgpu.GPUCommandBuffer{
			commandEncoder.Finish(),
		})
		return true
	})
	js.Global().Call("setInterval", updateGrid, updateInterval.Milliseconds())

	done := make(chan error, 1)
	<-done
}

func createQuadVertexData() []byte {
	var buffer bytes.Buffer
	writer := gblob.NewLittleEndianWriter(&buffer)
	// first vertex:
	writer.WriteFloat32(-0.9)
	writer.WriteFloat32(0.9)
	// second vertex:
	writer.WriteFloat32(-0.9)
	writer.WriteFloat32(-0.9)
	// third vertex:
	writer.WriteFloat32(0.9)
	writer.WriteFloat32(0.9)
	// fourth vertex:
	writer.WriteFloat32(0.9)
	writer.WriteFloat32(0.9)
	// fifth vertex:
	writer.WriteFloat32(-0.9)
	writer.WriteFloat32(-0.9)
	// sixth vertex:
	writer.WriteFloat32(0.9)
	writer.WriteFloat32(-0.9)
	return buffer.Bytes()
}

func createGridUniformData() []byte {
	var buffer bytes.Buffer
	writer := gblob.NewLittleEndianWriter(&buffer)
	writer.WriteFloat32(gridSize)
	writer.WriteFloat32(gridSize)
	return buffer.Bytes()
}

func createCellStateData() []byte {
	var buffer bytes.Buffer
	writer := gblob.NewLittleEndianWriter(&buffer)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < gridSize*gridSize; i++ {
		if rnd.Float64() > 0.7 {
			writer.WriteUint32(1)
		} else {
			writer.WriteUint32(0)
		}
	}
	return buffer.Bytes()
}

const renderShaderCode = `
struct VertexInput {
  @location(0) pos: vec2f,
  @builtin(instance_index) instance: u32,
};

struct VertexOutput {
  @builtin(position) pos: vec4f,
  @location(0) cell: vec2f,
};

@group(0) @binding(0) var<uniform> grid: vec2f;
@group(0) @binding(1) var<storage> cellState: array<u32>;

@vertex
fn vertexMain(input: VertexInput) -> VertexOutput {
	let i = f32(input.instance);
	let cell = vec2f(i % grid.x, floor(i / grid.x));
	let state = f32(cellState[input.instance]);

	let cellOffset = 2.0 * cell / grid;
	let gridPos = (input.pos * state + 1.0) / grid - 1.0 + cellOffset;

	var output: VertexOutput;
	output.pos = vec4f(gridPos, 0.0, 1.0);
	output.cell = cell;
	return output;
}

@fragment
fn fragmentMain(input: VertexOutput) -> @location(0) vec4f {
	let scaledCell = input.cell / grid;
	return vec4f(scaledCell, 1.0-scaledCell.x, 1.0);
}
`

const computeShaderCode = `
@group(0) @binding(0) var<uniform> grid: vec2f;

@group(0) @binding(1) var<storage> cellStateIn: array<u32>;
@group(0) @binding(2) var<storage, read_write> cellStateOut: array<u32>;

fn cellIndex(cell: vec2u) -> u32 {
	let clampX = cell.x % u32(grid.x);
	let clampY = cell.y % u32(grid.y);
  return clampX + clampY * u32(grid.x);
}

fn cellActive(x: u32, y: u32) -> u32 {
	return cellStateIn[cellIndex(vec2(x, y))];
}

@compute
@workgroup_size(8, 8)
fn computeMain(@builtin(global_invocation_id) cell: vec3u) {
	let countActive = cellActive(cell.x - 1u, cell.y + 0u) +
										cellActive(cell.x + 1u, cell.y + 0u) +
										cellActive(cell.x - 1u, cell.y + 1u) +
										cellActive(cell.x + 0u, cell.y + 1u) +
										cellActive(cell.x + 1u, cell.y + 1u) +
										cellActive(cell.x - 1u, cell.y - 1u) +
										cellActive(cell.x + 0u, cell.y - 1u) +
										cellActive(cell.x + 1u, cell.y - 1u);

	let cellSlot = cellIndex(cell.xy);
	switch countActive {
		case 2u: {
			cellStateOut[cellSlot] = cellStateIn[cellSlot];
		}
		case 3u: {
			cellStateOut[cellSlot] = 1u;
		}
		default: {
			cellStateOut[cellSlot] = 0u;
		}
	}
}
`
