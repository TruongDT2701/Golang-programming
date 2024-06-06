package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect the
	// "destroy" signal to the GtkMainQuit function to exit the GTK main
	// loop when the window is closed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Hello, GTK!")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new button widget with the label "Click Me!".
	btn, err := gtk.ButtonNewWithLabel("Click Me!")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	// Connect the "clicked" signal of the button to a function to be called
	// when it is clicked. This will cause the application to exit when
	// the button is clicked.
	btn.Connect("clicked", func() {
		log.Println("Button clicked!")
		gtk.MainQuit()
	})

	// Add the button to the window.
	win.Add(btn)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Run the GTK main loop. This blocks until gtk.MainQuit() is run.
	gtk.Main()
}
