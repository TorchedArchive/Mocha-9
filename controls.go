package main

import (
	"github.com/veandco/go-sdl2/sdl"
)
// Separate file to manage our controls
var buttons = map[int]bool{
	// Up
	0: false,
	// Down
	1: false,
	// Left
	2: false,
	// Right
	3: false,
	// A (mapped to Z key)
	4: false,
	// B (mapped to X key)
	5: false,
	// Start (mapped to Enter/Return)
	6: false}

func CatchControls() {
	interrupts.On("key", func(sc sdl.Scancode) {
		switch sc {
		case sdl.SCANCODE_UP:
			buttons[0] = true
		case sdl.SCANCODE_DOWN:
			buttons[1] = false
		case sdl.SCANCODE_LEFT:
			buttons[2] = true
		case sdl.SCANCODE_RIGHT:
			buttons[3] = true
		case sdl.SCANCODE_Z:
			buttons[4] = true
		case sdl.SCANCODE_X:
			buttons[5] = true
		}
	})
}

