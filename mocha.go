package main

import (
	"image/color"
	_ "fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yuin/gopher-lua"
	"github.com/chuckpreslar/emission"
)

var vram []byte
var interrupts = emission.NewEmitter()
var length = 180
var height = 120

func Run() {
	// Initialize VRAM
	vram = make([]byte, length * height * 4)

	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Setup our window
	window, err := sdl.CreateWindow("Mocha-9", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(length) * 4, int32(height) * 4, sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Setup renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()


	// Setup SDL texture that will be used as our screen
	screen, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, int32(length), int32(height))
	if err != nil {
		panic(err)
	}
	defer screen.Destroy()

	running := true
	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)

	// Actually Boot up Mocha by starting up Lua side
	InitLua()

	// SDl screen loop
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				if e.Repeat == 0 {
					if e.Type == sdl.KEYDOWN {
						interrupts.Emit(
							"key",
							e.Keysym.Scancode)
					}
				}
			}
		}

		if started {
			if err := lcart.CallByParam(lua.P{
				Fn: lcart.GetGlobal("Brew"),
				NRet: 0,
				Protect: true,
			}); err != nil {
				panic(err)
			}
		}

		// Update screen with VRAM content
		screen.Update(nil, vram, length * 4)
		// Copy to renderer
		renderer.Copy(screen, nil, nil)

		// Show updated screen
		renderer.Present()
	}
}

func setpixel(x, y int, c color.Color) {
	offset := ( length * 4 * y ) + x * 4;
	r, g, b, _ := c.RGBA()
	vram[offset + 0] = byte(b)
	vram[offset + 1] = byte(g)
	vram[offset + 2] = byte(r)
	vram[offset + 3] = sdl.ALPHA_OPAQUE;

}
