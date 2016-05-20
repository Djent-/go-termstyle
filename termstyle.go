package main

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/djent-/go-termstyle/ui"
	"regexp"
)

func main() {
	// Initialize gtk
	gtk.Init(nil)

	// Create MainWindow object
	mw := ui.NewMainWindow()
	mw.ShowAll()

	// Execute gtk main loop
	gtk.Main()
}

func parseXresources(data string) (colors []*gdk.RGBA, err error) {
	// Use regular expressions to exfiltrate hex values

	return
}

func connectDefines(data string) (colorMap []string) {
	// If the .Xresources file uses #defines everywhere,
	// I need to link the defined name to the hex value
	defineregex := regexp.MustCompile(`#define\s*(#[0-9a-fA-F]+)$`)
	defines := defineregex.FindAllStringSubmatch(data, -1)

	// Then, link the .color[0-15]s with the variable
	assignregex := regexp.MustCompile(`color\d+:\s*(\w*)`)
	assigns := assignregex.FindAllStringSubmatch(data, -1)

	// Finally, go back and link the .color[0-15] with the value

}

