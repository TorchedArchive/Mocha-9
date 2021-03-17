package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func Run() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Mocha-9", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 81 * 4, 81 * 4, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	screen, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, 81, 81)
	if err != nil {
		panic(err)
	}
	defer screen.Destroy()

	running := true
	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)

	InitLua()
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
