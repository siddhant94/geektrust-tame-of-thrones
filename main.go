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
