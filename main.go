package main

import (
	"colorstr/colorstr"
	"fmt"
)

func main() {
	text := "hi, finch. How you doing?"

	colortext1 := colorstr.Colorize([]string{"BrightMagentaBg"}, text)
	colortext2 := colorstr.Colorize([]string{"BrightMagentaFg"}, text)
	colortext3 := colorstr.Colorize([]string{"BlueFg", "MagentaBg"}, text)
	fmt.Println(colortext1)
	fmt.Println(colortext2)
	fmt.Println(colortext3)

	colortextrgb1 := colorstr.ColorizeRgbFg("#ffc8dd", text)
	colortextrgb2 := colorstr.ColorizeRgbBg("#780000", text)
	colortextrgb3 := colorstr.ColorizeRgb("#ffd60a", "#6d597a", text)
	fmt.Println(colortextrgb1)
	fmt.Println(colortextrgb2)
	fmt.Println(colortextrgb3)
}
