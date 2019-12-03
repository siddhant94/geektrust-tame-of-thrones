package main

import (
	"geektrust/utility"
	"fmt"
)

const inputFile = "input1.txt"

var messages []string

func main() {
	messages := utility.ReadInput(inputFile)
	fmt.Println(messages)
}
