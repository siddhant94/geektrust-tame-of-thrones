package messages

import "strings"

func checkAllegiance(msg string, emblem string) bool {
	var (
		emblemCharCount = make(map[int]int)
		actualCharCount = make(map[int]int)
	)
	for _, c := range emblem {
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

// Comparator takes 2 maps and checks for equality with condition being the two maps should have same keys and
// values of map1 should be lesser or equal to the map2.
func comparator(map1, map2 map[int]int) bool {
	// Compare both maps for length i.e. no of keys
	if len(map1) != len(map2) {
		return false
	}
	// Compare both maps for keys and values (condition: val of m1 <= val of m2); m1 = map1, m2= map2
	for key, val := range map2 {
		if _, ok := map1[key]; ok {
			if val < map1[key] {
				return false
			}

		} else {
			return false
		}
	}
	return true
}