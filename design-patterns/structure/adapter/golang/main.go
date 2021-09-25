package main

import (
	"fmt"
)

func main() {
	var notebook HDMI = &Notebook{"This is content."}
	fmt.Println(notebook.HDMIOut())

	// hdmi to vga
	var vga VGA = NewAdapterHDMIToVGA(notebook)
	fmt.Println(vga.VGAOut())
}
