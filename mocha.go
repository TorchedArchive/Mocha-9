package main

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

var vram []byte

func Run() {
	// Initialize VRAM
	vram = make([]byte, 144 * 81 * 4)

	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Setup our window
	window, err := sdl.CreateWindow("Mocha-9", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 144 * 4, 81 * 4, sdl.WINDOW_SHOWN)
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
	screen, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, 144, 81)
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
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		// Update screen with VRAM content
		screen.Update(nil, vram, 144 * 4)
		// Copy to renderer
		renderer.Copy(screen, nil, nil)

		// Show updated screen
		renderer.Present()
	}
}

func setpixel(x, y int, c color.Color) {
	offset := ( 144 * 4 * y ) + x * 4;
	r, g, b, _ := c.RGBA()
	vram[offset + 0] = byte(b)
	vram[offset + 1] = byte(g)
	vram[offset + 2] = byte(r)
	vram[offset + 3] = sdl.ALPHA_OPAQUE;

}
