package utility

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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
	//fmt.Println("Scan completed and reached EOF")
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	return result
}

// SplitMessages takes a slice of input messages and splits around each instance of one or more consecutive white space
// characters
func SplitMessages(messages []string) map[string]string {
	resultMap := make(map[string]string)
	for _, val := range messages {
		s := strings.Fields(val)
		resultMap[s[0]] = s[1]
	}
	return resultMap
}

// ProcessMessages
func ProcessMessages(inp map[string]string) []string {
	var allyKingdoms []string
	for kingdomName, msg := range inp {
		emblem := getEmblem(kingdomName)
		cipherKey := len(emblem)
		decryptedMsg := decryptMessage(msg, cipherKey)
		if checkAllegiance(decryptedMsg, emblem) {
			allyKingdoms = append(allyKingdoms, kingdomName)
		}
	}
	return allyKingdoms
}

func checkAllegiance(msg string, emblem string) bool {
	emblemCharCount := make(map[int]int)
	actualCharCount := make(map[int]int)
	for _, c := range emblem {
		//if !strings.Contains(strings.ToUpper(msg), string(c)) {
		//	return false
		//}
		if _, ok := emblemCharCount[int(c)]; ok {
			emblemCharCount[int(c)] += 1
		} else {
			emblemCharCount[int(c)] = 1
		}

		actualCharCount[int(c)] = strings.Count(msg, string(c))
	}
	res := comparator(emblemCharCount, actualCharCount)
	return res
}

// comparator takes 2 maps and checks for equality with condition being the two maps should have same keys and
// values of map1 should be lesser or equal to the map2.
//TODO: Find better and generic way to handle this.
func comparator(map1, map2 map[int]int) bool {
	// Compare both maps for length i.e. no of keys
	if len(map1) != len(map2) {
		return false
	}
	// Compare both maps for keys only
	for key, val := range map2 {
		if _, ok:= map1[key]; ok {
			if val < map1[key] {
				return false
			}
				
		} else {
			return false
		}
	}
	return true
}


func getEmblem(kingdomName string) string {
	switch kingdomName {
	case "LAND":
		return "PANDA"
	case "WATER":
		return "OCTOPUS"
	case "ICE":
		return "MAMMOTH"
	case "AIR":
		return "OWL"
	case "FIRE":
		return "DRAGON"
	default:
		return ""
	}
}

func decryptMessage(msg string, key int) string {
	var b strings.Builder
	b.Grow(len(msg))
	for _, val := range msg {
		asciiVal := int(val)
		decodedAscii := decipherFunc(asciiVal, key)
		fmt.Fprintf(&b, "%c", rune(decodedAscii))
	}
	s := b.String()
	s = s[:b.Len()]
	return s
}

func decipherFunc(asciiInt int, key int) int {
	if key > 26 {
		key = key % 26
	}
	// A - Z : 65-90 && a - z : 97 - 122
	diff := asciiInt - key
	if diff < 65 && !unicode.IsLower(rune(asciiInt)) {
		rotation := 64 - (diff) // Since 65 is counted
		return 90 - rotation
	}
	if diff < 97 && unicode.IsLower(rune(asciiInt)) {
		rotation := 96 - (diff) // Since 97 is counted
		return 122 - rotation
	}
	return diff
}