package ui

import (
	"fmt"
	"os"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/cairo"
	"github.com/djent-/go-termstyle/utils"
)

type CairoColor struct {
	R float64
	G float64
	B float64
}

type XresourcesFormat struct {
	Preamble       string
	Postamble      string
	DotColor       string
	DotForeground  string
	DotBackground  string
	DotCursorColor string
}

type MainWindow struct {
	// The Window
	*gtk.Window

	// Areas
	MainArea    *gtk.Box
	StyleArea   *gtk.Box
	PreviewArea *gtk.DrawingArea
	ColorArea   *gtk.Box
	MetaArea    *gtk.Box
	FileArea    *gtk.Box

	// Labels
	SpecialL *gtk.Label
	BlackL   *gtk.Label
	RedL     *gtk.Label
	GreenL   *gtk.Label
	YellowL  *gtk.Label
	BlueL    *gtk.Label
	MagentaL *gtk.Label
	CyanL    *gtk.Label
	WhiteL   *gtk.Label

	// Boxes
	SpecialB *gtk.Box
	BlackB   *gtk.Box
	RedB     *gtk.Box
	GreenB   *gtk.Box
	YellowB  *gtk.Box
	BlueB    *gtk.Box
	MagentaB *gtk.Box
	CyanB    *gtk.Box
	WhiteB   *gtk.Box

	// ColorButtons
	SpecialDark  *gtk.ColorButton // .background
	SpecialLight *gtk.ColorButton // .foreground and .cursorcolor
	BlackDark    *gtk.ColorButton
	BlackLight   *gtk.ColorButton
	RedDark      *gtk.ColorButton
	RedLight     *gtk.ColorButton
	GreenDark    *gtk.ColorButton
	GreenLight   *gtk.ColorButton
	YellowDark   *gtk.ColorButton
	YellowLight  *gtk.ColorButton
	BlueDark     *gtk.ColorButton
	BlueLight    *gtk.ColorButton
	MagentaDark  *gtk.ColorButton
	MagentaLight *gtk.ColorButton
	CyanDark     *gtk.ColorButton
	CyanLight    *gtk.ColorButton
	WhiteDark    *gtk.ColorButton
	WhiteLight   *gtk.ColorButton

	// Buttons
	ExportButton  *gtk.Button
	ImportButton  *gtk.Button
	RefreshButton *gtk.Button

	CurrentFormat *XresourcesFormat
}

func Newmainwindow() (mw *MainWindow) {
	// Create main window
	mw = new(MainWindow)
	// Create MainWindow Window
	mw.Window, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	mw.Window.SetTitle("TermStyle")
	mw.Window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	mw.Window.SetDefaultSize(650, 700)
	// Create MainArea
	mw.MainArea, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	// Create StyleArea
	mw.StyleArea, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	// Create ExportArea
	mw.MetaArea, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	// Create DrawArea
	mw.PreviewArea, _ = gtk.DrawingAreaNew()
	mw.PreviewArea.Connect("draw", mw.draw)
	// Create ColorArea
	mw.ColorArea, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	// Create FileArea
	mw.FileArea, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 3)

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
	sdRGBA := gdk.NewRGBA(29.0/256, 31.0/256, 33.0/256, 1)
	mw.SpecialDark.SetRGBA(sdRGBA)
	mw.SpecialLight.SetUseAlpha(false)
	slRGBA := gdk.NewRGBA(197.0/256, 200.0/256, 198.0/256, 1)
	mw.SpecialLight.SetRGBA(slRGBA)

	mw.BlackDark.SetUseAlpha(false)
	bkdRGBA := gdk.NewRGBA(40.0/256, 42.0/256, 46.0/256, 1)
	mw.BlackDark.SetRGBA(bkdRGBA)
	mw.BlackLight.SetUseAlpha(false)
	bklRGBA := gdk.NewRGBA(55.0/256, 59.0/256, 65.0/256, 1)
	mw.BlackLight.SetRGBA(bklRGBA)

	mw.RedDark.SetUseAlpha(false)
	rdRGBA := gdk.NewRGBA(165.0/256, 66.0/256, 66.0/256, 1)
	mw.RedDark.SetRGBA(rdRGBA)
	mw.RedLight.SetUseAlpha(false)
	rlRGBA := gdk.NewRGBA(204.0/256, 102.0/256, 102.0/256, 1)
	mw.RedLight.SetRGBA(rlRGBA)

	mw.GreenDark.SetUseAlpha(false)
	gdRGBA := gdk.NewRGBA(140.0/256, 148.0/256, 64.0/256, 1)
	mw.GreenDark.SetRGBA(gdRGBA)
	mw.GreenLight.SetUseAlpha(false)
	glRGBA := gdk.NewRGBA(181.0/256, 189.0/256, 104.0/256, 1)
	mw.GreenLight.SetRGBA(glRGBA)

	mw.YellowDark.SetUseAlpha(false)
	ydRGBA := gdk.NewRGBA(222.0/256, 147.0/256, 95.0/256, 1)
	mw.YellowDark.SetRGBA(ydRGBA)
	mw.YellowLight.SetUseAlpha(false)
	ylRGBA := gdk.NewRGBA(240.0/256, 198.0/256, 116.0/256, 1)
	mw.YellowLight.SetRGBA(ylRGBA)

	mw.BlueDark.SetUseAlpha(false)
	bdRGBA := gdk.NewRGBA(95.0/256, 129.0/256, 157.0/256, 1)
	mw.BlueDark.SetRGBA(bdRGBA)
	mw.BlueLight.SetUseAlpha(false)
	blRGBA := gdk.NewRGBA(129.0/256, 162.0/256, 190.0/256, 1)
	mw.BlueLight.SetRGBA(blRGBA)

	mw.MagentaDark.SetUseAlpha(false)
	mdRGBA := gdk.NewRGBA(133.0/256, 103.0/256, 143.0/256, 1)
	mw.MagentaDark.SetRGBA(mdRGBA)
	mw.MagentaLight.SetUseAlpha(false)
	mlRGBA := gdk.NewRGBA(179.0/256, 148.0/256, 187.0/256, 1)
	mw.MagentaLight.SetRGBA(mlRGBA)

	mw.CyanDark.SetUseAlpha(false)
	cdRGBA := gdk.NewRGBA(94.0/256, 141.0/256, 135.0/256, 1)
	mw.CyanDark.SetRGBA(cdRGBA)
	mw.CyanLight.SetUseAlpha(false)
	clRGBA := gdk.NewRGBA(138.0/256, 190.0/256, 183.0/256, 1)
	mw.CyanLight.SetRGBA(clRGBA)

	mw.WhiteDark.SetUseAlpha(false)
	wdRGBA := gdk.NewRGBA(112.0/256, 120.0/256, 128.0/256, 1)
	mw.WhiteDark.SetRGBA(wdRGBA)
	mw.WhiteLight.SetUseAlpha(false)
	wlRGBA := gdk.NewRGBA(197.0/256, 200.0/256, 198.0/256, 1)
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
	mw.RefreshButton, _ = gtk.ButtonNewWithLabel("Refresh")
	mw.RefreshButton.Connect("clicked", func() {
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

	// Create export button
	mw.ExportButton, _ = gtk.ButtonNewWithLabel("Export .Xresources")
	mw.ExportButton.Connect("clicked", mw.saveDialog)

	// Create import button
	mw.ImportButton, _ = gtk.ButtonNewWithLabel("Import .Xresources")
	mw.ImportButton.Connect("clicked", mw.openDialog)

	// Pack FileArea
	mw.FileArea.PackStart(mw.ImportButton, true, true, 1)
	mw.FileArea.Add(mw.ExportButton)

	// Pack MetaArea
	mw.MetaArea.PackStart(mw.FileArea, true, true, 1)
	mw.MetaArea.Add(mw.RefreshButton)

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
			cr.Rectangle(float64(x-1)*w*2.3, float64(y-1)*h, w, h)
			cr.SetSourceRGB(colors[z].R, colors[z].G, colors[z].B)
			cr.Fill()
			cr.MoveTo(float64(x-1)*w*2.3+w+5, float64(y-1)*h+h/2)
			cr.SetFontSize(fontsize)
			cr.ShowText(hexvals[z])
			cr.Fill()
			z++
		}
	}
}

func (mw *MainWindow) getColors() (colors []*gdk.RGBA) {
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

func (mw *MainWindow) saveDialog(button *gtk.Button) {
	// Create a FileChooserDialog
	filechooser, _ := gtk.FileChooserDialogNewWith2Buttons(
		"Save As",                    // Dialog title
		mw.Window,                    // Parent Window
		gtk.FILE_CHOOSER_ACTION_SAVE, // File Chooser Action
		"Cancel",                     // Button 1 Text
		gtk.RESPONSE_CANCEL,          // Response Type
		"Save",                       // Button 2 Text
		gtk.RESPONSE_OK,              // Response Type
	)

	// Set the default filename as ".Xresources"
	filechooser.SetCurrentName(".Xresources")
	// Get a response
	response := filechooser.Run()
	// Get the information
	switch response {
	case -5: // case gtk.RESPONSE_OK
		filename := filechooser.GetFilename()
		filechooser.Destroy()
		mw.exportAs(filename)
	case -6: // case gtk.RESPONSE_CANCEL
		filechooser.Destroy()
	}
}

func (mw *MainWindow) exportAs(filename string) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		// File exists
		// Spawn an alert dialog asking whether to overwrite

		// If dialog comes back OK, overwrite, else return nil
		return
	} else {
		// Create and write to filename
		f, err := os.Create(filename)
		if err != nil {
			// Spawn an error message - no permissions or disk full probably
			panic(err)
		}
		defer f.Close()

		// Generate the .Xresources file format with the color data
		f.WriteString(mw.exportString())
		f.Sync()
	}
}

func (mw *MainWindow) openDialog(button *gtk.Button) {
	// Create a FileChooserDialog
	filechooser, _ = gtk.FileChooserDialogNewWith2Buttons(
		"Open",                       // Dialog title
		mw.Window,                    // Parent Window
		gtk.FILE_CHOOSER_ACTION_OPEN, // File Chooser Action
		"Cancel",                     // Button 1 text
		gtk.RESPONSE_CANCEL,          // Response type
		"Open",                       // Button 2 text
		gtk.RESPONSE_OK,              // Response type
	)

	// Set the default filename as ".Xresources"
	filechooser.SetCurrentName(".Xresources")
	// Get a response
	response := filechooser.Run()
	// Get the information
	switch response {
	case -5: // case gtk.RESPONSE_OK
		filename := filechooser.GetFilename()
		filechooser.Destroy()
		mw.importXresources(filename)
	case -6: // case gtk.RESPONSE_CANCEL
		filechooser.Destroy()
	}
}

func (mw *MainWindow) importXresources(filename string) {
	// Parse pre-existing .Xresources
	data, err := ioutil.ReadFile(filename)
	if err != nil {

		return
	}
	// Open file

	// Assign new values to colorbuttons and refresh
	colors, err := parseXresources(string(data))
	mw.assignColorButtons(colors)
}

var dotcolor string = "URxvt*color"
var dotforeground string = "URxvt.foreground"
var dotbackground string = "URxvt.background"
var dotcursorcolor string = "URxvt.cursorColor"
var preamble string = `Xft.dpi: 180
URxvt.scrollBar: false
URxvt.font: xft:dejavu sans mono:size=10
! URxvt.letterSpace: -3`
var postamble string = "! vim: ft=xdefaults"

func (mw *MainWindow) exportString() (export string) {
	colors := convertAllRGBAtoHex(mw.getColors())
	export = fmt.Sprintf("%s\n", preamble)
	export = export + fmt.Sprintf("! special\n%s: %s\n", dotforeground, colors[1])
	export = export + fmt.Sprintf("%s: %s\n", dotbackground, colors[0])
	export = export + fmt.Sprintf("%s: %s\n\n", dotcursorcolor, colors[1])
	export = export + fmt.Sprintf("! black\n%s0: %s\n", dotcolor, colors[2])
	export = export + fmt.Sprintf("%s8: %s\n\n", dotcolor, colors[3])
	export = export + fmt.Sprintf("! red\n%s1: %s\n", dotcolor, colors[4])
	export = export + fmt.Sprintf("%s9: %s\n\n", dotcolor, colors[5])
	export = export + fmt.Sprintf("! green\n%s2: %s\n", dotcolor, colors[6])
	export = export + fmt.Sprintf("%s10: %s\n\n", dotcolor, colors[7])
	export = export + fmt.Sprintf("! yellow\n%s3: %s\n", dotcolor, colors[8])
	export = export + fmt.Sprintf("%s11: %s\n\n", dotcolor, colors[9])
	export = export + fmt.Sprintf("! blue\n%s4: %s\n", dotcolor, colors[10])
	export = export + fmt.Sprintf("%s12: %s\n\n", dotcolor, colors[11])
	export = export + fmt.Sprintf("! magenta\n%s5: %s\n", dotcolor, colors[12])
	export = export + fmt.Sprintf("%s13: %s\n\n", dotcolor, colors[13])
	export = export + fmt.Sprintf("! cyan\n%s6: %s\n", dotcolor, colors[14])
	export = export + fmt.Sprintf("%s14: %s\n\n", dotcolor, colors[15])
	export = export + fmt.Sprintf("! white\n%s7: %s\n", dotcolor, colors[16])
	export = export + fmt.Sprintf("%s15: %s\n\n", dotcolor, colors[17])
	export = export + postamble
	return
}

func (mw *MainWindow) assignColorButtons(colors []*gtk.RGBA) {
	// Assign new colors to colorbuttons

	return
}
