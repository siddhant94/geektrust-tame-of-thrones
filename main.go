package main

import (
	"fmt"
	"geektrust/utility"
	"os"
)

var messages []string

func main() {
	//var inputFile = "input1.txt"
	if len(os.Args) <= 1 {
		fmt.Println("Provide Input File path")
		return
	}
	inputFile := os.Args[1] // Default File path is dropped for user provided path
	messages := utility.ReadInput(inputFile)
	kingdomMsgMap := utility.SplitMessages(messages)
	allies := utility.ProcessMessages(kingdomMsgMap)
	if len(allies) < 3 {
		fmt.Println("NONE")
	} else {
		res := "SPACE "
		for _, val := range allies {
			res += val+", "
		}
		res = res[:len(res)-2] // Remove last trailing comma and space
		fmt.Println(res)
	}
}
