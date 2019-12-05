package utility

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		//fmt.Println(scanner.Text())
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

// ProcessMessages takes a slice of encrypted messages and decrypts it
func DecryptMessages(messages []string) map[string]string {
	resultMap := make(map[string]string)
	// const msgRegex = `(?i)(?P<Name>Air|Land|Water|Ice|Fire),(?:\s*)"(?P<Message>.*?)"`
	const msgRegex = `(?i)(?P<Name>Air|Land|Water|Ice|Fire)(?:\s)(?P<Message>.*)`
	r, err := regexp.Compile(msgRegex)
	if err != nil {
		fmt.Printf("Failed to compile Regex.", err)
		return resultMap
	}

	for _, val := range messages {
		//matches := r.FindAllString(val, -1)
		match := r.FindStringSubmatch(val)
		for i, name := range r.SubexpNames() {
			if i > 0 && i <= len(match) {
				resultMap[name] = match[i]
			}
		}
	}
	return resultMap
}