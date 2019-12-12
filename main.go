package main

import (
	"fmt"
	"geektrust/utility"
	"os"
)

func main() {
	var inputFile string
	if len(os.Args) <= 1 {
		fmt.Println("File Path not provided in arguments. Using default Path : " + inputFile)
		inputFile = "input1.txt"
	} else {
		inputFile = os.Args[1]
	}
	var messages = utility.ReadInput(inputFile)
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
