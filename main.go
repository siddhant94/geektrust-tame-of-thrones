package main

import (
	"fmt"
	"geektrust/kingdoms"
	"geektrust/msgprocess"
	"os"
)

var DefaultInputFile string
func init() {
	DefaultInputFile = "input1.txt"
}

func main() {
	// Create Message sender Kingdom, for the problem in question it is King Shan's
	senderKingdom := kingdoms.NewKingdom("SPACE", "GORILLA")

	var messageSrc string
	if len(os.Args) <= 1 {
		messageSrc = DefaultInputFile
		fmt.Println("File Path not provided in arguments. Using default File : " + messageSrc)
	} else {
		messageSrc = os.Args[1]
	}
	var messagesList = msgprocess.ReadFileInput(messageSrc)

	// Check for empty file
	if(len(messagesList) < 1) {
		fmt.Println("Input file is empty...terminating")
		return
	}

	// Parse the Input data into map{Msg Recipient(kingdom) : Message}
	kingdomToMsgMap := msgprocess.ParseInputData(messagesList)

	// Process the messages to determine allies by checking msg success.
	allies := msgprocess.ProcessMessages(kingdomToMsgMap)
	if len(allies) < 3 {
		fmt.Println("NONE")
	} else {
		res := senderKingdom.Name + " "
		for _, val := range allies {
			res += val + ", "
		}
		res = res[:len(res)-2] // Remove last trailing comma and space
		fmt.Println(res)
	}
}
