package main

import (
	"bufio"
	"os"

	"github.com/anilvdeshpande/GoLangToDoClientApp/menu"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	//Display the option Menu
	menu.ShowMenu()

	// Capture the option
	menu.CaptureOptionFromMenu()

}
