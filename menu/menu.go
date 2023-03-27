package menu

import "fmt"

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

func ShowMenu() {
	fmt.Println("Select one of the options below")
	for i := 0; i < len(menuOptions); i++ {
		fmt.Printf("%v. %v\n", (i + 1), menuOptions[i])
	}
}
