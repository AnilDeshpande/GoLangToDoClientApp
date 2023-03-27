package menu

import (
	"bufio"
	"fmt"
	"os"
)

var MenuOptions = []string{
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
	for i := 0; i < len(MenuOptions); i++ {
		fmt.Printf("%v. %v\n", (i + 1), MenuOptions[i])
	}
}
