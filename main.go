package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFile = "input1.txt"
var messages []string

func main() {
	messages := readInput(inputFile)
}

func readInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
}