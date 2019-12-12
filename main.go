package main

import (
	"fmt"
	"geektrust/utility"
	"os"
)

func main() {
	var inputFile string
	if len(os.Args) <= 1 {
		inputFile = "input1.txt"
		fmt.Println("File Path not provided in arguments. Using default File : " + inputFile)
	} else {
		inputFile = os.Args[1]
	}
	var messages = utility.ReadInput(inputFile)
	if(len(messages) < 1) {
		return
	}
	kingdomMsgMap := utility.SplitMessages(messages)
	allies := utility.ProcessMessages(kingdomMsgMap)
	if len(allies) < 3 {
		fmt.Println("NONE")
	} else {
		res := "SPACE "
		for _, val := range allies {
			res += val + ", "
		}
		res = res[:len(res)-2] // Remove last trailing comma and space
		fmt.Println(res)
	}
}
