package colorstr

// the color format information comes from https://en.wikipedia.org/wiki/ANSI_escape_code

import (
	"fmt"
	"strconv"
	"strings"
)

var Color = map[string]string{
	"BlackFg":         "30",
	"RedFg":           "31",
	"GreenFg":         "32",
	"YellowFg":        "33",
	"BlueFg":          "34",
	"MagentaFg":       "35",
	"CyanFg":          "36",
	"WhiteFg":         "37",
	"BrightBlackFg":   "90",
	"BrightRedFg":     "91",
	"BrightGreenFg":   "92",
	"BrightYellowFg":  "93",
	"BrightBlueFg":    "94",
	"BrightMagentaFg": "95",
	"BrightCyanFg":    "96",
	"BrightWhiteFg":   "97",
	"BlackBg":         "40",
	"RedBg":           "41",
	"GreenBg":         "42",
	"YellowBg":        "43",
	"BlueBg":          "44",
	"MagentaBg":       "45",
	"CyanBg":          "46",
	"WhiteBg":         "47",
	"BrightBlackBg":   "100",
	"BrightRedBg":     "101",
	"BrightGreenBg":   "102",
	"BrightYellowBg":  "103",
	"BrightBlueBg":    "104",
	"BrightMagentaBg": "105",
	"BrightCyanBg":    "106",
	"BrightWhiteBg":   "107",
	"End":             "0",
}

// create a type to receive command args
type ColorPair struct {
	Fg string
	Bg string
}

// function to init a colorpair type
func NewColorPair(fg, bg string) *ColorPair {
	return &ColorPair{
		Fg: fg,
		Bg: bg,
	}
}

// can receive more than one color, foreground color and background color
// cause color format depends on the color number, so the order is not important
// \033[30;45m is the same as \033[45;30m
// colorname can be []string{foregrondcolor backgroundcolor} or []string{backgroundcolor foregroundcolor}
func Colorize(colorname []string, text string) string {
	var colortext string

	colorcode1, _ := Color[colorname[0]]
	colortext = fmt.Sprintf("\033[%sm%s\033[0m", colorcode1, text)

	if len(colorname) == 2 {
		colorcode2, _ := Color[colorname[1]]
		colortext = fmt.Sprintf("\033[%s;%sm%s\033[0m", colorcode1, colorcode2, text)
	}

	return colortext
}

func ColorizeRgbFg(rgb, text string) string {

	r, g, b := rgb[1:3], rgb[3:5], rgb[5:7]
	numr, _ := strconv.ParseUint(r, 16, 8)
	numg, _ := strconv.ParseUint(g, 16, 8)
	numb, _ := strconv.ParseUint(b, 16, 8)

	colorizeText := fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", numr, numg, numb, text)
	return colorizeText
}

func ColorizeRgbBg(rgb, text string) string {

	r, g, b := rgb[1:3], rgb[3:5], rgb[5:7]
	numr, _ := strconv.ParseUint(r, 16, 8)
	numg, _ := strconv.ParseUint(g, 16, 8)
	numb, _ := strconv.ParseUint(b, 16, 8)

	colorizeText := fmt.Sprintf("\033[48;2;%d;%d;%dm%s\033[0m", numr, numg, numb, text)
	return colorizeText
}

// can receive forground color and background color but the order in important
// rgb1 = foreground color, rgb2 = background color
func ColorizeRgb(foreground, background, text string) string {

	r1, g1, b1 := foreground[1:3], foreground[3:5], foreground[5:7]
	numr1, _ := strconv.ParseUint(r1, 16, 8)
	numg1, _ := strconv.ParseUint(g1, 16, 8)
	numb1, _ := strconv.ParseUint(b1, 16, 8)

	r2, g2, b2 := background[1:3], background[3:5], background[5:7]
	numr2, _ := strconv.ParseUint(r2, 16, 8)
	numg2, _ := strconv.ParseUint(g2, 16, 8)
	numb2, _ := strconv.ParseUint(b2, 16, 8)

	colorizeText := fmt.Sprintf("\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm%s\033[0m", numr1, numg1, numb1, numr2, numg2, numb2, text)
	return colorizeText
}

// check terminal args color mode, used for automatically RenderText()
func checkMode(c *ColorPair) string {
	var mode string

	if strings.HasPrefix(c.Fg, "#") || strings.HasPrefix(c.Bg, "#") {
		mode = "rgb"
	} else {
		mode = "ascll"
	}

	return mode
}

// check whether the color is predefined or not
func checkAscll(c []string) bool {
	for _, j := range c {
		if _, exists := Color[j]; !exists {
			fmt.Printf("color: %s is not right", j)
			return false
		}
	}

	return true
}

// check whether the color is correct rgb or not
func checkRgb(c []string) bool {
	for _, j := range c {
		if len(j) != 7 || !strings.HasPrefix(j, "#") {
			fmt.Printf("color: %s is not right", j)
			return false
		}
	}

	return true
}

// used in uncertain situation, when the color type and color count is uncertain
// a automatical way to check color exists and color type
func RenderText(c *ColorPair, text string) string {
	var mode = checkMode(c)
	var colortext string

	switch mode {
	case "rgb":
		if c.Fg != "nil" && c.Bg != "nil" && checkRgb([]string{c.Fg, c.Bg}) {
			colortext = ColorizeRgb(c.Fg, c.Bg, text)
		}
		if c.Fg != "nil" && c.Bg == "nil" && checkRgb([]string{c.Fg}) {
			colortext = ColorizeRgbFg(c.Fg, text)
		}
		if c.Fg == "nil" && c.Bg != "nil" && checkRgb([]string{c.Bg}) {
			colortext = ColorizeRgbBg(c.Bg, text)
		}
	case "ascll":
		if c.Fg != "nil" && c.Bg != "nil" && checkAscll([]string{c.Fg, c.Bg}) {
			colortext = Colorize([]string{c.Fg, c.Bg}, text)
		}
		if c.Fg != "nil" && c.Bg == "nil" && checkAscll([]string{c.Fg}) {
			colortext = Colorize([]string{c.Fg}, text)
		}
		if c.Fg == "nil" && c.Bg != "nil" && checkAscll([]string{c.Bg}) {
			colortext = Colorize([]string{c.Bg}, text)
		}
	}

	return colortext
}
