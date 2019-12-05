package main

import (
	"fmt"
	"geektrust/utility"
)

const inputFile = "input1.txt"

var messages []string

func main() {
	// TODO: Make input file to be read from os.Args
	messages := utility.ReadInput(inputFile)
	fmt.Println(messages)
	resultMap := utility.DecryptMessages(messages)
	fmt.Println(resultMap)

}
