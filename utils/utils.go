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
	hexadecimal = fmt.Sprintf("#%s%s%s", HexEnc(r), HexEnc(g), HexEnc(b))
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

func HextoRGBA(hexval string) (rgba *gdk.RGBA) {
	// Input: "#123abc"
	hexdigits := make(map[string]int)
	hdigits := [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	for i, v := range hdigits {
		hexdigits[v] = i
	}
	// Split hexval string into the three bytes
	// Multiply the top four bits by 16 then add the bottom four bits
	byte1 := hexval[1:3]
	byte2 := hexval[3:5]
	byte3 := hexval[5:]
	r := float64(hexdigits[string(byte1[0])]*16 + hexdigits[string(byte1[1])])
	g := float64(hexdigits[string(byte2[0])]*16 + hexdigits[string(byte2[1])])
	b := float64(hexdigits[string(byte3[0])]*16 + hexdigits[string(byte3[1])])

	rgba = gdk.NewRGBA(r/float64(256), g/float64(256), b/float64(256), 1)
	return
}

func ConvertAllHextoRGBA(hexvals map[string]string) (rgbas map[string]*gdk.RGBA) {
	rgbas = make(map[string]*gdk.RGBA)
	for color, hex := range hexvals {
		rgbas[color] = HextoRGBA(hex)
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
