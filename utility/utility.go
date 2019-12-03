package utility

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	var result []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) // Tokenizes based on new line
	for scanner.Scan() {           // Value is false when EOF or error
		fmt.Println(scanner.Text())
		result = append(result, scanner.Text())
	}

	err = scanner.Err()
	if err == nil {
		fmt.Println("Scan completed and reached EOF")
	} else {
		fmt.Println(err)
		return []string{}
	}
	return result
}