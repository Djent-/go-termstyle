package ui

import (
	"fmt"
	"github.com/djent-/go-termstyle/utils"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

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
	RefreshArea *gtk.Box

	// Labels
	SpecialL    *gtk.Label
	BlackL      *gtk.Label
	RedL        *gtk.Label
	GreenL      *gtk.Label
	YellowL     *gtk.Label
	BlueL       *gtk.Label
	MagentaL    *gtk.Label
	CyanL       *gtk.Label
	WhiteL      *gtk.Label
	EyedropperL *gtk.Label

	colorButtonBoxPadding uint
	colorAreaPadding      uint

	EyedropperImage *gtk.Image
	ColorComboBox   *gtk.ComboBoxText
	EyedropperColor string

	// Boxes
	SpecialB      *gtk.Box
	BlackB        *gtk.Box
	RedB          *gtk.Box
	GreenB        *gtk.Box
	YellowB       *gtk.Box
	BlueB         *gtk.Box
	MagentaB      *gtk.Box
	CyanB         *gtk.Box
	WhiteB        *gtk.Box
	EyedropperBox *gtk.Box

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
	ExportButton       *gtk.Button
	ImportButton       *gtk.Button
	RefreshButton      *gtk.Button
	EyedropperButton   *gtk.Button
	EyedropperOKButton *gtk.Button
	CurrentFormat      *XresourcesFormat
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
	mw.Window.SetDefaultSize(650, 700)
	// Create MainArea
	mw.MainArea, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	// Create StyleArea
	mw.StyleArea, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	// Create MetaArea
	mw.MetaArea, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	// Create DrawArea
	mw.PreviewArea, _ = gtk.DrawingAreaNew()
	mw.PreviewArea.Connect("draw", mw.draw)
	// Create ColorArea
	mw.ColorArea, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	// Create FileArea
	mw.FileArea, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 3)

	// Set default Xresources format
	mw.CurrentFormat = &XresourcesDefault

	// Pack Areas
	mw.MainArea.PackStart(mw.StyleArea, false, false, 3)
	mw.MainArea.Add(mw.MetaArea)
	mw.StyleArea.PackStart(mw.PreviewArea, true, true, 5)
	mw.StyleArea.Add(mw.ColorArea)
	mw.RefreshArea, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)

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
	sdRGBA := gdk.NewRGBA(29.0/255, 31.0/255, 33.0/255, 1)
	mw.SpecialDark.SetRGBA(sdRGBA)
	mw.SpecialLight.SetUseAlpha(false)
	slRGBA := gdk.NewRGBA(197.0/255, 200.0/255, 198.0/255, 1)
	mw.SpecialLight.SetRGBA(slRGBA)

	mw.BlackDark.SetUseAlpha(false)
	bkdRGBA := gdk.NewRGBA(40.0/255, 42.0/255, 46.0/255, 1)
	mw.BlackDark.SetRGBA(bkdRGBA)
	mw.BlackLight.SetUseAlpha(false)
	bklRGBA := gdk.NewRGBA(55.0/255, 59.0/255, 65.0/255, 1)
	mw.BlackLight.SetRGBA(bklRGBA)

	mw.RedDark.SetUseAlpha(false)
	rdRGBA := gdk.NewRGBA(165.0/255, 66.0/255, 66.0/255, 1)
	mw.RedDark.SetRGBA(rdRGBA)
	mw.RedLight.SetUseAlpha(false)
	rlRGBA := gdk.NewRGBA(204.0/255, 102.0/255, 102.0/255, 1)
	mw.RedLight.SetRGBA(rlRGBA)

	mw.GreenDark.SetUseAlpha(false)
	gdRGBA := gdk.NewRGBA(140.0/255, 148.0/255, 64.0/255, 1)
	mw.GreenDark.SetRGBA(gdRGBA)
	mw.GreenLight.SetUseAlpha(false)
	glRGBA := gdk.NewRGBA(181.0/255, 189.0/255, 104.0/255, 1)
	mw.GreenLight.SetRGBA(glRGBA)

	mw.YellowDark.SetUseAlpha(false)
	ydRGBA := gdk.NewRGBA(222.0/255, 147.0/255, 95.0/255, 1)
	mw.YellowDark.SetRGBA(ydRGBA)
	mw.YellowLight.SetUseAlpha(false)
	ylRGBA := gdk.NewRGBA(240.0/255, 198.0/255, 116.0/255, 1)
	mw.YellowLight.SetRGBA(ylRGBA)

	mw.BlueDark.SetUseAlpha(false)
	bdRGBA := gdk.NewRGBA(95.0/255, 129.0/255, 157.0/255, 1)
	mw.BlueDark.SetRGBA(bdRGBA)
	mw.BlueLight.SetUseAlpha(false)
	blRGBA := gdk.NewRGBA(129.0/255, 162.0/255, 190.0/255, 1)
	mw.BlueLight.SetRGBA(blRGBA)

	mw.MagentaDark.SetUseAlpha(false)
	mdRGBA := gdk.NewRGBA(133.0/255, 103.0/255, 143.0/255, 1)
	mw.MagentaDark.SetRGBA(mdRGBA)
	mw.MagentaLight.SetUseAlpha(false)
	mlRGBA := gdk.NewRGBA(179.0/255, 148.0/255, 187.0/255, 1)
	mw.MagentaLight.SetRGBA(mlRGBA)

	mw.CyanDark.SetUseAlpha(false)
	cdRGBA := gdk.NewRGBA(94.0/255, 141.0/255, 135.0/255, 1)
	mw.CyanDark.SetRGBA(cdRGBA)
	mw.CyanLight.SetUseAlpha(false)
	clRGBA := gdk.NewRGBA(138.0/255, 190.0/255, 183.0/255, 1)
	mw.CyanLight.SetRGBA(clRGBA)

	mw.WhiteDark.SetUseAlpha(false)
	wdRGBA := gdk.NewRGBA(112.0/255, 120.0/255, 128.0/255, 1)
	mw.WhiteDark.SetRGBA(wdRGBA)
	mw.WhiteLight.SetUseAlpha(false)
	wlRGBA := gdk.NewRGBA(197.0/255, 200.0/255, 198.0/255, 1)
	mw.WhiteLight.SetRGBA(wlRGBA)

	// Pack ColorButton Boxes
	mw.colorButtonBoxPadding = 0
	mw.SpecialB.PackStart(mw.SpecialDark, true, true, mw.colorButtonBoxPadding)
	mw.SpecialB.Add(mw.SpecialLight)
	mw.BlackB.PackStart(mw.BlackDark, true, true, mw.colorButtonBoxPadding)
	mw.BlackB.Add(mw.BlackLight)
	mw.RedB.PackStart(mw.RedDark, true, true, mw.colorButtonBoxPadding)
	mw.RedB.Add(mw.RedLight)
	mw.GreenB.PackStart(mw.GreenDark, true, true, mw.colorButtonBoxPadding)
	mw.GreenB.Add(mw.GreenLight)
	mw.YellowB.PackStart(mw.YellowDark, true, true, mw.colorButtonBoxPadding)
	mw.YellowB.Add(mw.YellowLight)
	mw.BlueB.PackStart(mw.BlueDark, true, true, mw.colorButtonBoxPadding)
	mw.BlueB.Add(mw.BlueLight)
	mw.MagentaB.PackStart(mw.MagentaDark, true, true, mw.colorButtonBoxPadding)
	mw.MagentaB.Add(mw.MagentaLight)
	mw.CyanB.PackStart(mw.CyanDark, true, true, mw.colorButtonBoxPadding)
	mw.CyanB.Add(mw.CyanLight)
	mw.WhiteB.PackStart(mw.WhiteDark, true, true, mw.colorButtonBoxPadding)
	mw.WhiteB.Add(mw.WhiteLight)

	// Create redraw button
	mw.RefreshButton, _ = gtk.ButtonNewWithLabel("Refresh")
	mw.RefreshButton.Connect("clicked", func() {
		mw.PreviewArea.QueueDraw()
	})

	// Pack ColorArea
	mw.colorAreaPadding = 1
	mw.ColorArea.PackStart(mw.SpecialL, true, true, mw.colorAreaPadding)
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
	mw.ExportButton.Connect("clicked", mw.exportXresources)

	// Create import button
	mw.ImportButton, _ = gtk.ButtonNewWithLabel("Import .Xresources")
	mw.ImportButton.Connect("clicked", mw.importXresources)

	// Create EyedropperButton and add it to its box
	mw.EyedropperButton, _ = gtk.ButtonNew()
	//"home/watermelon/go/src/github.com/djent-/go-termstyle/resources/eyedropper-icon-32x32.png"
	EyedropperImage, err := gtk.ImageNewFromFile("./resources/eyedropper-icon-32x32.png")
	if err != nil {
		panic(err)
	}
	mw.EyedropperImage = EyedropperImage
	// DesignContest: www.designcontest.com CC Attribution 4.0
	mw.EyedropperButton.SetImage(mw.EyedropperImage)
	mw.EyedropperButton.SetAlwaysShowImage(true)
	mw.EyedropperButton.Connect("clicked", mw.getEyedropperColor)
	mw.EyedropperBox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	mw.EyedropperBox.PackStart(mw.EyedropperButton, false, false, 1)

	// Create EyedropperL
	mw.EyedropperL, _ = gtk.LabelNew("#FFFFFF")
	mw.EyedropperColor = "#FFFFFF"
	mw.setEyedropperLMarkup()

	// Create EyedropperOKButton
	mw.EyedropperOKButton, _ = gtk.ButtonNewWithLabel("OK")
	// TODO: setEyedropperColor
	mw.EyedropperOKButton.Connect("clicked", mw.setEyedropperColor)

	// Create ColorSwitcher
	mw.ColorComboBox, _ = gtk.ComboBoxTextNew()
	mw.setColorComboBoxEntriesR()

	// Pack FileArea
	mw.FileArea.PackStart(mw.ImportButton, true, true, 1)
	mw.FileArea.PackEnd(mw.ExportButton, true, true, 1)

	// Pack RefreshArea
	mw.RefreshArea.PackStart(mw.EyedropperBox, false, false, 5)
	mw.RefreshArea.Add(mw.EyedropperL)
	mw.RefreshArea.Add(mw.ColorComboBox)
	mw.RefreshArea.Add(mw.EyedropperOKButton)
	mw.RefreshArea.PackEnd(mw.RefreshButton, true, true, 5)

	// Pack MetaArea
	mw.MetaArea.PackStart(mw.FileArea, true, true, 1)
	mw.MetaArea.Add(mw.RefreshArea)

	// Add MainArea to the Window
	mw.Window.Add(mw.MainArea)

	// Return mw
	return
}

func (mw *MainWindow) draw(da *gtk.DrawingArea, cr *cairo.Context) {
	// Get values from all the ColorButtons
	colors := utils.ConvertAllRGBAtoCC(mw.getColors())
	hexvals := utils.ConvertAllRGBAtoHex(mw.getColors())
	// Draw two columns of rectangles
	h, w := float64(85), float64(100)
	// Loop
	z := 0
	var fontsize float64 = 25.0
	for y := 1; y <= 9; y++ {
		for x := 1; x <= 2; x++ {
			//log.Println("draw z:", z)
			cr.Rectangle(float64(x-1)*w*2.3, float64(y-1)*h+5, w, h)
			cr.SetSourceRGB(colors[z].R, colors[z].G, colors[z].B)
			cr.Fill()
			cr.MoveTo(float64(x-1)*w*2.3+w+5, float64(y-1)*h+h/2+5)
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

func (mw *MainWindow) saveDialog() (filename string) {
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
		filename = filechooser.GetFilename()
		filechooser.Destroy()
	case -6: // case gtk.RESPONSE_CANCEL
		filechooser.Destroy()
	}
	return
}

func (mw *MainWindow) exportXresources() {
	filename := mw.saveDialog()
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

func (mw *MainWindow) openDialog() (filename string) {
	// Create a FileChooserDialog
	filechooser, _ := gtk.FileChooserDialogNewWith2Buttons(
		"Open",                       // Dialog title
		mw.Window,                    // Parent Window
		gtk.FILE_CHOOSER_ACTION_OPEN, // File Chooser Action
		"Cancel",                     // Button 1 text
		gtk.RESPONSE_CANCEL,          // Response type
		"Open",                       // Button 2 text
		gtk.RESPONSE_OK,              // Response type
	)

	// Get a response
	response := filechooser.Run()
	// Get the information
	switch response {
	case -5: // case gtk.RESPONSE_OK
		filename = filechooser.GetFilename()
		filechooser.Destroy()
	case -6: // case gtk.RESPONSE_CANCEL
		filechooser.Destroy()
	}
	return
}

func (mw *MainWindow) importXresources() {
	// Parse pre-existing .Xresources
	filename := mw.openDialog()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	// Assign new values to colorbuttons and refresh
	colors, err := parseXresources(string(data))
	mw.assignColorButtons(colors)
	mw.PreviewArea.QueueDraw()
}

var (
	XresourcesDefault = XresourcesFormat{
		DotColor:       ".color",
		DotForeground:  ".foreground",
		DotBackground:  ".background",
		DotCursorColor: ".cursorColor",
		Preamble:       "",
		Postamble:      "! vim: ft=xdefaults"}
	XresourcesURxvtHiDPINoScrollBar = XresourcesFormat{
		DotColor:       "URxvt*color",
		DotForeground:  "URxvt*foreground",
		DotBackground:  "URxvt*background",
		DotCursorColor: "URxvt*cursorColor",
		Preamble: `Xft.dpi: 180
URxvt.scrollBar: false
URxvt.font: xft:dejavu sans mono:size=10
! URxvt.letterSpace: -3`,
		Postamble: "! vim: ft=xdefaults"}
)

func (mw *MainWindow) exportString() (export string) {
	colors := utils.ConvertAllRGBAtoHex(mw.getColors())
	f := mw.CurrentFormat
	export = fmt.Sprintf("%s\n", f.Preamble)
	export = export + fmt.Sprintf("! special\n%s: %s\n", f.DotForeground, colors[1])
	export = export + fmt.Sprintf("%s: %s\n", f.DotBackground, colors[0])
	export = export + fmt.Sprintf("%s: %s\n\n", f.DotCursorColor, colors[1])
	export = export + fmt.Sprintf("! black\n%s0: %s\n", f.DotColor, colors[2])
	export = export + fmt.Sprintf("%s8: %s\n\n", f.DotColor, colors[3])
	export = export + fmt.Sprintf("! red\n%s1: %s\n", f.DotColor, colors[4])
	export = export + fmt.Sprintf("%s9: %s\n\n", f.DotColor, colors[5])
	export = export + fmt.Sprintf("! green\n%s2: %s\n", f.DotColor, colors[6])
	export = export + fmt.Sprintf("%s10: %s\n\n", f.DotColor, colors[7])
	export = export + fmt.Sprintf("! yellow\n%s3: %s\n", f.DotColor, colors[8])
	export = export + fmt.Sprintf("%s11: %s\n\n", f.DotColor, colors[9])
	export = export + fmt.Sprintf("! blue\n%s4: %s\n", f.DotColor, colors[10])
	export = export + fmt.Sprintf("%s12: %s\n\n", f.DotColor, colors[11])
	export = export + fmt.Sprintf("! magenta\n%s5: %s\n", f.DotColor, colors[12])
	export = export + fmt.Sprintf("%s13: %s\n\n", f.DotColor, colors[13])
	export = export + fmt.Sprintf("! cyan\n%s6: %s\n", f.DotColor, colors[14])
	export = export + fmt.Sprintf("%s14: %s\n\n", f.DotColor, colors[15])
	export = export + fmt.Sprintf("! white\n%s7: %s\n", f.DotColor, colors[16])
	export = export + fmt.Sprintf("%s15: %s\n\n", f.DotColor, colors[17])
	export = export + f.Postamble
	return
}

func (mw *MainWindow) assignColorButtons(colors map[string]*gdk.RGBA) {
	// Assign new colors to colorbuttons
	mw.SpecialDark.SetRGBA(colors["background"])
	mw.SpecialLight.SetRGBA(colors["foreground"])
	mw.BlackDark.SetRGBA(colors["color0"])
	mw.BlackLight.SetRGBA(colors["color8"])
	mw.RedDark.SetRGBA(colors["color1"])
	mw.RedLight.SetRGBA(colors["color9"])
	mw.GreenDark.SetRGBA(colors["color2"])
	mw.GreenLight.SetRGBA(colors["color10"])
	mw.YellowDark.SetRGBA(colors["color3"])
	mw.YellowLight.SetRGBA(colors["color11"])
	mw.BlueDark.SetRGBA(colors["color4"])
	mw.BlueLight.SetRGBA(colors["color12"])
	mw.MagentaDark.SetRGBA(colors["color5"])
	mw.MagentaLight.SetRGBA(colors["color13"])
	mw.CyanDark.SetRGBA(colors["color6"])
	mw.CyanLight.SetRGBA(colors["color14"])
	mw.WhiteDark.SetRGBA(colors["color7"])
	mw.WhiteLight.SetRGBA(colors["color15"])
	return
}

func (mw *MainWindow) getEyedropperColor() {
	// Execute grabc if installed, capture output, and set label markup
	out, _ := exec.Command("grabc").Output()
	mw.setEyedropperLText(string(out[:7]))
	return
}

func (mw *MainWindow) setEyedropperColor() {
	// Called when the OK button is pressed
	// Sets the color of the colorbutton corresponding to the combobox
	// Create an array of colorbuttons
	colorbuttons := [18]*gtk.ColorButton{mw.SpecialDark, mw.SpecialLight, mw.BlackDark, mw.RedDark, mw.GreenDark, mw.YellowDark, mw.BlueDark, mw.MagentaDark, mw.CyanDark, mw.WhiteDark, mw.BlackLight, mw.RedLight, mw.GreenLight, mw.YellowLight, mw.BlueLight, mw.MagentaLight, mw.CyanLight, mw.WhiteLight}
	// Get a map of formatted color names to ints
	colornames := utils.ArraySwap(mw.getFormattedColorNames())
	colorbuttons[colornames[mw.ColorComboBox.GetActiveText()]].SetRGBA(utils.HextoRGBA(mw.EyedropperColor))
	return
}

func (mw *MainWindow) setColorComboBoxEntriesF() {
	// Take the current XresourcesFormat and populate the entries
	// of the ColorComboBox dropdown menu
	// First empty the ComboBox
	mw.ColorComboBox.RemoveAll()
	colornames := mw.getFormattedColorNames()
	for _, v := range colornames {
		mw.ColorComboBox.AppendText(v)
	}
	return
}

func (mw *MainWindow) setColorComboBoxEntriesR() {
	// Populate the entries in the ColorComboBox based on the true
	// names of the colors
	mw.ColorComboBox.RemoveAll()
	colornames := [18]string{"SpecialDark", "SpecialLight", "BlackDark", "RedDark", "GreenDark", "YellowDark", "BlueDark", "MagentaDark", "CyanDark", "WhiteDark", "BlackLight", "RedLight", "GreenLight", "YellowLight", "BlueLight", "MagentaLight", "CyanLight", "WhiteLight"}
	for _, v := range colornames {
		mw.ColorComboBox.AppendText(v)
	}
	return
}

func (mw *MainWindow) getFormattedColorNames() (colornames [18]string) {
	colornames[0] = mw.CurrentFormat.DotBackground
	colornames[1] = mw.CurrentFormat.DotForeground
	for i := 0; i <= 15; i++ {
		colornames[i+2] = mw.CurrentFormat.DotColor + strconv.Itoa(i)
	}
	return
}

func (mw *MainWindow) setEyedropperLText(text string) {
	mw.EyedropperColor = strings.ToUpper(text)
	mw.setEyedropperLMarkup()
	return
}

func (mw *MainWindow) setEyedropperLMarkup() {
	// Set the text color of the label to the hex value in its text
	color := mw.EyedropperColor
	markup := fmt.Sprintf("<span color=\"%s\">%s</span>", color, color)
	mw.EyedropperL.SetMarkup(markup)
	return
}

func parseXresources(data string) (colors map[string]*gdk.RGBA, err error) {
	// If the Xresources contains defines
	if matched, _ := regexp.Match(`#define`, []byte(data)); matched {
		colors = utils.ConvertAllHextoRGBA(connectDefines(data))
	} else {
		// else just go through normally
		assignregex := regexp.MustCompile(`(color\d*|foreground|background|cursorColor):\s*(#[0-9a-fA-F]+)`)
		assigns := assignregex.FindAllStringSubmatch(data, -1)
		assignmap := make(map[string]string)
		for _, match := range assigns {
			assignmap[match[1]] = match[2]
		}
		colors = utils.ConvertAllHextoRGBA(assignmap)
	}
	return
}

func connectDefines(data string) (colorMap map[string]string) {
	// If the .Xresources file uses #defines everywhere,
	// I need to link the defined name to the hex value
	defineregex := regexp.MustCompile(`#define\s*(\w*)\s*(#[0-9a-fA-F]+)`)
	defines := defineregex.FindAllStringSubmatch(data, -1)

	// Then, link the .color[0-15]s with the variable
	assignregex := regexp.MustCompile(`(color\d*|foreground|background|cursorColor):\s*(\w*)`)
	assigns := assignregex.FindAllStringSubmatch(data, -1)

	// Finally, go back and link the .color[0-15] with the value
	// Data is in the format [["string matched" "parens1" "parens2"]]
	// I could make this more compact, but it would be less readable
	definemap := make(map[string]string)
	for _, match := range defines {
		definemap[match[1]] = match[2]
	}
	assignmap := make(map[string]string)
	for _, match := range assigns {
		assignmap[match[1]] = match[2]
	}
	for color, assignment := range assignmap {
		colorMap[color] = definemap[assignment]
	}
	return
}
