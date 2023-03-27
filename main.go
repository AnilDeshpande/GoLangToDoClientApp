package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"menu/menu.go"
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

func main() {
	//Display the option Menu

	// Capture the option
	selection := ReadInteger()
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

	//Perform the action
	// Ask for continuation

}

func ReadInteger() int8 {
	input, err := reader.ReadString('\n')
	HandleError(err)
	selection, err := strconv.ParseInt(strings.TrimSpace(input), 0, 8)
	HandleError(err)
	return int8(selection)
}

func ReadString() string {
	input, err := reader.ReadString('\n')
	HandleError(err)
	return strings.TrimSpace(input)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
