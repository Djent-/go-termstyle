package utils

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"strconv"
	"strings"
)

type CairoColor struct {
	R float64
	G float64
	B float64
}

func RGBAtoCairoColor(color *gdk.RGBA) (cc CairoColor) {
	// Convert RGBA (0-256, 0-256, 0-256) to float64 (0-1, 0-1, 0-1) for cairo
	floats := color.Floats()
	cc.R = floats[0]
	cc.G = floats[1]
	cc.B = floats[2]
	return
}

func ConvertAllRGBAtoCC(colors []*gdk.RGBA) (ccs []CairoColor) {
	// Convert a slice of gdk.RGBA colors to CairoColors
	for _, rgba := range colors {
		ccs = append(ccs, RGBAtoCairoColor(rgba))
	}
	return
}

func RGBAtoHex(color *gdk.RGBA) (hexadecimal string) {
	rgbastring := color.String()
	rgbastring = strings.Replace(rgbastring, "rgb(", "", -1)
	rgbastring = strings.Replace(rgbastring, ")", "", -1)
	rgbvals := strings.Split(rgbastring, ",")
	r, _ := strconv.Atoi(rgbvals[0])
	g, _ := strconv.Atoi(rgbvals[1])
	b, _ := strconv.Atoi(rgbvals[2])
	hexadecimal = fmt.Sprintf("#%s%s%s", hexenc(r), hexenc(g), hexenc(b))
	//log.Println("Added ", hexadecimal)
	return
}

func ConvertAllRGBAtoHex(colors []*gdk.RGBA) (hexcolors []string) {
	// Convert a slice of gdk.RGBA colors to the format #00bbFF
	for _, rgba := range colors {
		hexcolors = append(hexcolors, RGBAtoHex(rgba))
	}
	return
}

func HexEnc(i int) (hexadecimal string) {
	hexdigits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	dig1 := 0
	for ; i >= 16; i = i - 16 {
		dig1++
	}
	return fmt.Sprintf("%s%s", hexdigits[dig1], hexdigits[i])
}
