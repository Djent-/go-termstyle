package main

import (
	"github.com/djent-/go-termstyle/ui"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
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
