package main

import (
	"fmt"
	"os"
	"image/color"

	"github.com/PinwheelSystem/PaletteNom"
	"github.com/PinwheelSystem/bitmap"
	"github.com/akamensky/argparse"
)

var font map[string]bitmap.Glyph
var palette []color.Color

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
	palette, err = palettelib.Load("palettes/monobrew.png")

	if err != nil {
		panic(err)
	}

	Run()
}
