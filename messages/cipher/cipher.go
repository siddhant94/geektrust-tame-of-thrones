package cipher

import (
	"fmt"
	"strings"
	"unicode"
)

func DecryptMessage(msg string, key int) string {
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
