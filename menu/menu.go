package menu

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anilvdeshpande/GoLangToDoClientApp/utils"
)

var menuOptions = []string{
	"Register User",
	"Login User",
	"Add ToDoItem",
	"Remove ToDoItem",
	"Update ToDoItem",
	"Remove ToDoItem",
	"Logout User",
	"Exit",
}

var reader = bufio.NewReader(os.Stdin)

func ShowMenu() {
	fmt.Println("Select one of the options below")
	for i := 0; i < len(menuOptions); i++ {
		fmt.Printf("%v. %v\n", (i + 1), menuOptions[i])
	}
}

func CaptureOptionFromMenu() {
	selection := utils.ReadInteger()
	if selection < 1 || selection > 8 {
		panic("Wrong Selection: Exiting")
	} else {
		fmt.Println("Selected Option: ", menuOptions[selection-1])
		switch selection {
		case 1:
			fmt.Println("Performing ", menuOptions[selection-1])
		case 2:
			fmt.Println("Performing ", menuOptions[selection-1])
		case 3:
			fmt.Println("Performing ", menuOptions[selection-1])
		case 4:
			fmt.Println("Performing ", menuOptions[selection-1])
		case 5:
			fmt.Println("Performing ", menuOptions[selection-1])
		case 6:
			fmt.Println("Performing ", menuOptions[selection-1])
		case 7:
			fmt.Println("Performing ", menuOptions[selection-1])
		case 8:
			fmt.Println("Performing  ", menuOptions[selection-1])
			panic("Exiting")
		default:
			panic(" Wrong selection: Exiting")
		}
	}
}
