package msgprocess

import (
	"bufio"
	"fmt"
	"geektrust/kingdoms"
	"geektrust/msgprocess/cipher"
	"os"
	"strings"
)

func ReadFileInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	var result []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) // Tokenizes based on new line
	for scanner.Scan() {           // Value is false when EOF or error
		result = append(result, scanner.Text())
	}

	err = scanner.Err()
	//fmt.Println("Scan completed and reached EOF")
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	return result
}

// ParseInputData takes a slice of input messages and splits around each instance of one or more consecutive white space
// characters
func ParseInputData(messages []string) map[string]string {
	resultMap := make(map[string]string)
	for _, val := range messages {
		s := strings.Fields(val) // s = [ICE MAMMOTH THVAO]; s[0] = kingdom name, s[]
		kingdomName := s[0]
		s[0] = s[len(s)-1]
		s = s[:len(s)-1]
		msg := strings.Join(s, " ")
		resultMap[kingdomName] = msg
	}
	return resultMap
}


// ProcessMessages takes a map{kingdom:msg} sent by King Shan, decodes it, calls checkAllegience and returns with ally
// kingdoms
func ProcessMessages(inp map[string]string) []string {
	var allyKingdoms []string
	for kingdomName, msg := range inp {
		emblem := kingdoms.GetKindomData(kingdomName).Emblem
		cipherKey := len(emblem)
		decryptedMsg := cipher.DecryptMessage(msg, cipherKey)
		if checkAllegiance(decryptedMsg, emblem) {
			allyKingdoms = append(allyKingdoms, kingdomName)
		}
	}
	return allyKingdoms
}