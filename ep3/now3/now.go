package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"vladimirvivien/time/lib"
)

func main() {
	img(fmt.Sprintf("Current time: %s", lib.ShortTime()))
}

func img(str string) {
	const s = 1024
	dc := gg.NewContext(s, s)

	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	if err := dc.LoadFontFace("/Library/Fonts/Arial Unicode.ttf", 96); err != nil {
		panic(err)
	}

	dc.DrawStringAnchored(str, s/2, s/2, 0.5, 0.5)
	dc.SavePNG("out.png")
}
