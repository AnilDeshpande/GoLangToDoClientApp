package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

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
