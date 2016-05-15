package main

import(
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/cairo"
	"fmt"
	"strings"
	"strconv"
	"log"
)

type MainWindow struct {
	// The Window
	*gtk.Window

	// Areas
	MainArea *gtk.Box
	StyleArea *gtk.Box
	PreviewArea *gtk.DrawingArea
	ColorArea *gtk.Box
	ExportArea *gtk.Box

	// Labels
	SpecialL *gtk.Label
	BlackL *gtk.Label
	RedL *gtk.Label
	GreenL *gtk.Label
	YellowL *gtk.Label
	BlueL *gtk.Label
	MagentaL *gtk.Label
	CyanL *gtk.Label
	WhiteL *gtk.Label

	// Boxes
	SpecialB *gtk.Box
	BlackB *gtk.Box
	RedB *gtk.Box
	GreenB *gtk.Box
	YellowB *gtk.Box
	BlueB *gtk.Box
	MagentaB *gtk.Box
	CyanB *gtk.Box
	WhiteB *gtk.Box

	// ColorButtons
	SpecialDark *gtk.ColorButton // .background
	SpecialLight *gtk.ColorButton // .foreground and .cursorcolor
	BlackDark *gtk.ColorButton
	BlackLight *gtk.ColorButton
	RedDark *gtk.ColorButton
	RedLight *gtk.ColorButton
	GreenDark *gtk.ColorButton
	GreenLight *gtk.ColorButton
	YellowDark *gtk.ColorButton
	YellowLight *gtk.ColorButton
	BlueDark *gtk.ColorButton
	BlueLight *gtk.ColorButton
	MagentaDark *gtk.ColorButton
	MagentaLight *gtk.ColorButton
	CyanDark *gtk.ColorButton
	CyanLight *gtk.ColorButton
	WhiteDark *gtk.ColorButton
	WhiteLight *gtk.ColorButton

	// Buttons
	ExportButton *gtk.Button
	RedrawButton *gtk.Button

	// Dialogs
	//ExportDialog *gtk.FileChooserDialog
}

type CairoColor struct {
	R float64
	G float64
	B float64
}

func main() {
	// Initialize gtk
	gtk.Init(nil)

	// Create MainWindow object
	mw := NewMainWindow()
	mw.ShowAll()

	// Execute gtk main loop
	gtk.Main()
}

func NewMainWindow() (mw *MainWindow) {
	// Create main window
	mw = new(MainWindow)
	// Create MainWindow Window
	mw.Window, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	mw.Window.SetTitle("TermStyle")
	mw.Window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	mw.Window.SetDefaultSize(650,700)
	// Create MainArea
	mw.MainArea, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	// Create StyleArea
	mw.StyleArea, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	// Create ExportArea
	mw.ExportArea, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	// Create DrawArea
	mw.PreviewArea, _ = gtk.DrawingAreaNew()
	mw.PreviewArea.Connect("draw", mw.draw)
	// Create ColorArea
	mw.ColorArea, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)

	// Pack Areas
	mw.MainArea.PackStart(mw.StyleArea, false, false, 3)
	mw.MainArea.Add(mw.ExportArea)
	mw.StyleArea.PackStart(mw.PreviewArea, true, true, 5)
	mw.StyleArea.Add(mw.ColorArea)

	// Create labels
	// 9 in total
	mw.SpecialL, _ = gtk.LabelNew("Special")
	mw.BlackL, _ = gtk.LabelNew("Black")
	mw.RedL, _ = gtk.LabelNew("Red")
	mw.GreenL, _ = gtk.LabelNew("Green")
	mw.YellowL, _ = gtk.LabelNew("Yellow")
	mw.BlueL, _ = gtk.LabelNew("Blue")
	mw.MagentaL, _ = gtk.LabelNew("Magenta")
	mw.CyanL, _ = gtk.LabelNew("Cyan")
	mw.WhiteL, _ = gtk.LabelNew("White")

	// Create ColorButton Boxes
	mw.SpecialB, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	mw.BlackB, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	mw.RedB, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	mw.GreenB, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	mw.YellowB, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	mw.BlueB, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	mw.MagentaB, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	mw.CyanB, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	mw.WhiteB, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)

	// Create ColorButtons
	// 18 in total
	mw.SpecialDark, _ = gtk.ColorButtonNew()
	mw.SpecialLight, _ = gtk.ColorButtonNew()
	mw.BlackDark, _ = gtk.ColorButtonNew()
	mw.BlackLight, _ = gtk.ColorButtonNew()
	mw.RedDark, _ = gtk.ColorButtonNew()
	mw.RedLight, _ = gtk.ColorButtonNew()
	mw.GreenDark, _ = gtk.ColorButtonNew()
	mw.GreenLight, _ = gtk.ColorButtonNew()
	mw.YellowDark, _ = gtk.ColorButtonNew()
	mw.YellowLight, _ = gtk.ColorButtonNew()
	mw.BlueDark, _ = gtk.ColorButtonNew()
	mw.BlueLight, _ = gtk.ColorButtonNew()
	mw.MagentaDark, _ = gtk.ColorButtonNew()
	mw.MagentaLight, _ = gtk.ColorButtonNew()
	mw.CyanDark, _ = gtk.ColorButtonNew()
	mw.CyanLight, _ = gtk.ColorButtonNew()
	mw.WhiteDark, _ = gtk.ColorButtonNew()
	mw.WhiteLight, _ = gtk.ColorButtonNew()

	// Set ColorButtons to default state
	mw.SpecialDark.SetUseAlpha(false)
	sdRGBA := gdk.NewRGBA(29.0/256,31.0/256,33.0/256,1)
	mw.SpecialDark.SetRGBA(sdRGBA)
	mw.SpecialLight.SetUseAlpha(false)
	slRGBA := gdk.NewRGBA(197.0/256,200.0/256,198.0/256,1)
	mw.SpecialLight.SetRGBA(slRGBA)

	mw.BlackDark.SetUseAlpha(false)
	bkdRGBA := gdk.NewRGBA(40.0/256,42.0/256,46.0/256,1)
	mw.BlackDark.SetRGBA(bkdRGBA)
	mw.BlackLight.SetUseAlpha(false)
	bklRGBA := gdk.NewRGBA(55.0/256,59.0/256,65.0/256,1)
	mw.BlackLight.SetRGBA(bklRGBA)

	mw.RedDark.SetUseAlpha(false)
	rdRGBA := gdk.NewRGBA(165.0/256,66.0/256,66.0/256,1)
	mw.RedDark.SetRGBA(rdRGBA)
	mw.RedLight.SetUseAlpha(false)
	rlRGBA := gdk.NewRGBA(204.0/256,102.0/256,102.0/256,1)
	mw.RedLight.SetRGBA(rlRGBA)

	mw.GreenDark.SetUseAlpha(false)
	gdRGBA := gdk.NewRGBA(140.0/256,148.0/256,64.0/256,1)
	mw.GreenDark.SetRGBA(gdRGBA)
	mw.GreenLight.SetUseAlpha(false)
	glRGBA := gdk.NewRGBA(181.0/256,189.0/256,104.0/256,1)
	mw.GreenLight.SetRGBA(glRGBA)

	mw.YellowDark.SetUseAlpha(false)
	ydRGBA := gdk.NewRGBA(222.0/256,147.0/256,95.0/256,1)
	mw.YellowDark.SetRGBA(ydRGBA)
	mw.YellowLight.SetUseAlpha(false)
	ylRGBA := gdk.NewRGBA(240.0/256,198.0/256,116.0/256,1)
	mw.YellowLight.SetRGBA(ylRGBA)

	mw.BlueDark.SetUseAlpha(false)
	bdRGBA := gdk.NewRGBA(95.0/256,129.0/256,157.0/256,1)
	mw.BlueDark.SetRGBA(bdRGBA)
	mw.BlueLight.SetUseAlpha(false)
	blRGBA := gdk.NewRGBA(129.0/256,162.0/256,190.0/256,1)
	mw.BlueLight.SetRGBA(blRGBA)

	mw.MagentaDark.SetUseAlpha(false)
	mdRGBA := gdk.NewRGBA(133.0/256,103.0/256,143.0/256,1)
	mw.MagentaDark.SetRGBA(mdRGBA)
	mw.MagentaLight.SetUseAlpha(false)
	mlRGBA := gdk.NewRGBA(179.0/256,148.0/256,187.0/256,1)
	mw.MagentaLight.SetRGBA(mlRGBA)

	mw.CyanDark.SetUseAlpha(false)
	cdRGBA := gdk.NewRGBA(94.0/256,141.0/256,135.0/256,1)
	mw.CyanDark.SetRGBA(cdRGBA)
	mw.CyanLight.SetUseAlpha(false)
	clRGBA := gdk.NewRGBA(138.0/256,190.0/256,183.0/256,1)
	mw.CyanLight.SetRGBA(clRGBA)

	mw.WhiteDark.SetUseAlpha(false)
	wdRGBA := gdk.NewRGBA(112.0/256,120.0/256,128.0/256,1)
	mw.WhiteDark.SetRGBA(wdRGBA)
	mw.WhiteLight.SetUseAlpha(false)
	wlRGBA := gdk.NewRGBA(197.0/256,200.0/256,198.0/256,1)
	mw.WhiteLight.SetRGBA(wlRGBA)

	// Pack ColorButton Boxes
	mw.SpecialB.PackStart(mw.SpecialDark, true, true, 1)
	mw.SpecialB.Add(mw.SpecialLight)
	mw.BlackB.PackStart(mw.BlackDark, true, true, 1)
	mw.BlackB.Add(mw.BlackLight)
	mw.RedB.PackStart(mw.RedDark, true, true, 1)
	mw.RedB.Add(mw.RedLight)
	mw.GreenB.PackStart(mw.GreenDark, true, true, 1)
	mw.GreenB.Add(mw.GreenLight)
	mw.YellowB.PackStart(mw.YellowDark, true, true, 1)
	mw.YellowB.Add(mw.YellowLight)
	mw.BlueB.PackStart(mw.BlueDark, true, true, 1)
	mw.BlueB.Add(mw.BlueLight)
	mw.MagentaB.PackStart(mw.MagentaDark, true, true, 1)
	mw.MagentaB.Add(mw.MagentaLight)
	mw.CyanB.PackStart(mw.CyanDark, true, true, 1)
	mw.CyanB.Add(mw.CyanLight)
	mw.WhiteB.PackStart(mw.WhiteDark, true, true, 1)
	mw.WhiteB.Add(mw.WhiteLight)

	// Create redraw button
	mw.RedrawButton, _ = gtk.ButtonNewWithLabel("Refresh")
	mw.RedrawButton.Connect("clicked", func() {
		mw.PreviewArea.QueueDraw()
	})

	// Pack ColorArea
	mw.ColorArea.PackStart(mw.SpecialL, true, true, 3)
	mw.ColorArea.Add(mw.SpecialB)
	mw.ColorArea.Add(mw.BlackL)
	mw.ColorArea.Add(mw.BlackB)
	mw.ColorArea.Add(mw.RedL)
	mw.ColorArea.Add(mw.RedB)
	mw.ColorArea.Add(mw.GreenL)
	mw.ColorArea.Add(mw.GreenB)
	mw.ColorArea.Add(mw.YellowL)
	mw.ColorArea.Add(mw.YellowB)
	mw.ColorArea.Add(mw.BlueL)
	mw.ColorArea.Add(mw.BlueB)
	mw.ColorArea.Add(mw.MagentaL)
	mw.ColorArea.Add(mw.MagentaB)
	mw.ColorArea.Add(mw.CyanL)
	mw.ColorArea.Add(mw.CyanB)
	mw.ColorArea.Add(mw.WhiteL)
	mw.ColorArea.Add(mw.WhiteB)
	mw.ColorArea.Add(mw.RedrawButton)

	// Create export button
	mw.ExportButton, _ = gtk.ButtonNewWithLabel("Export .Xresources")
	mw.ExportButton.Connect("clicked", mw.saveFile)

	// Create FileChooserDialog
	//mw.ExportDialog, _ = gtk.FileChooserDialogNew()

	// Pack ExportArea
	mw.ExportArea.PackStart(mw.ExportButton, true, true, 1)

	// Add MainArea to the Window
	mw.Window.Add(mw.MainArea)

	// Return mw
	return
}

func (mw *MainWindow) draw(da *gtk.DrawingArea, cr *cairo.Context) {
	// Get values from all the ColorButtons
	colors := convertAllRGBAtoCC(mw.getColors())
	hexvals := convertAllRGBAtoHex(mw.getColors())
	// Draw two columns of rectangles
	h, w := float64(85), float64(100)
	// Loop
	z := 0
	var fontsize float64 = 25.0
	for y := 1; y <= 9; y++ {
		for x := 1; x <= 2; x++ {
			//log.Println("draw z:", z)
			cr.Rectangle(float64(x-1) * w * 2.3, float64(y-1) * h, w, h)
			cr.SetSourceRGB(colors[z].R, colors[z].G, colors[z].B)
			cr.Fill()
			cr.MoveTo(float64(x-1) * w * 2.3 + w + 5, float64(y-1) * h + h / 2)
			cr.SetFontSize(fontsize)
			cr.ShowText(hexvals[z])
			cr.Fill()
			z++
		}
	}
}

func (mw *MainWindow) getColors() (colors []*gdk.RGBA){
	colors = append(colors, mw.SpecialDark.GetRGBA())
	colors = append(colors, mw.SpecialLight.GetRGBA())
	colors = append(colors, mw.BlackDark.GetRGBA())
	colors = append(colors, mw.BlackLight.GetRGBA())
	colors = append(colors, mw.RedDark.GetRGBA())
	colors = append(colors, mw.RedLight.GetRGBA())
	colors = append(colors, mw.GreenDark.GetRGBA())
	colors = append(colors, mw.GreenLight.GetRGBA())
	colors = append(colors, mw.YellowDark.GetRGBA())
	colors = append(colors, mw.YellowLight.GetRGBA())
	colors = append(colors, mw.BlueDark.GetRGBA())
	colors = append(colors, mw.BlueLight.GetRGBA())
	colors = append(colors, mw.MagentaDark.GetRGBA())
	colors = append(colors, mw.MagentaLight.GetRGBA())
	colors = append(colors, mw.CyanDark.GetRGBA())
	colors = append(colors, mw.CyanLight.GetRGBA())
	colors = append(colors, mw.WhiteDark.GetRGBA())
	colors = append(colors, mw.WhiteLight.GetRGBA())
	return
}

func (mw *MainWindow) saveFile(button *gtk.Button) {
	// Create a FileChooserDialog
	filechooser, _ := gtk.FileChooserDialogNewWith2Buttons(
		"Save As", // Dialog title
		mw.Window, // Parent Window
		gtk.FILE_CHOOSER_ACTION_SAVE, // File Chooser Action
		"Cancel", // Button 1 Text
		gtk.RESPONSE_CANCEL, // Response Type
		"Save", // Button 2 Text
		gtk.RESPONSE_OK,  // Response Type
	)
	// Show the FileChooserDialog
	filechooser.Show()

	// Get the information
	folder, _ := filechooser.GetCurrentFolder()
	filename := filechooser.GetFilename()
	log.Println("Folder: ", folder)
	log.Println("Filename:", filename)
}

func RGBAtoCairoColor(color *gdk.RGBA) (cc CairoColor) {
	// Convert RGBA (0-256, 0-256, 0-256) to float64 (0-1, 0-1, 0-1) for cairo
	floats := color.Floats()
	cc.R = floats[0]
	cc.G = floats[1]
	cc.B = floats[2]
	return
}

func convertAllRGBAtoCC(colors []*gdk.RGBA) (ccs []CairoColor) {
	// Convert a slice of gdk.RGBA colors to CairoColors
	for _, rgba := range colors {
		ccs = append(ccs, RGBAtoCairoColor(rgba))
		//log.Println("colors []*gdk.RGBA index:", index)
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
	//log.Println("r, g, b: ", r, g, b)
	hexadecimal = fmt.Sprintf("#%s%s%s", hexenc(r), hexenc(g), hexenc(b))
	//log.Println("Added ", hexadecimal)
	return
}

func convertAllRGBAtoHex(colors []*gdk.RGBA) (hexcolors []string) {
	// Convert a slice of gdk.RGBA colors to the format #00bbFF
	for _, rgba := range colors {
		hexcolors = append(hexcolors, RGBAtoHex(rgba))
	}
	return
}

func hexenc(i int) (hexadecimal string) {
	hexdigits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	dig1 := 0
	for ; i >= 16; i = i - 16{
		dig1++
	}
	return fmt.Sprintf("%s%s", hexdigits[dig1], hexdigits[i])
}
