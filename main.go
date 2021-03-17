package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var font map[string]bitmap.Glyph

func main() {
	parser := argparse.NewParser("mocha-9", "A sweet warm fantasy brew.")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	
	// Load bitmap font
	// TODO: Use a different font
	bm := bitmap.New()
	font = bm.Load("m5x7.png")

	// Load palette for
	// TODO: Monochrome palette, like gameboy screen (green and green-black)
	palettelib := palettenom.New()
	colors, err := palettelib.Load("coffeebrew.png")
	if err != nil {
		panic(err)
	}

	Run()
}
