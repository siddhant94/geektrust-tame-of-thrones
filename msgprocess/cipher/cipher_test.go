package cipher

import "testing"

func TestDecipherFunc(t *testing.T) {
	var testCases = []struct {
		Input      int
		CipherKey  int
		DecodedRes int
	}{
		{70, 5, 65},    // (70) F ,key(5) => A, Ascii Val = 65
		{82, 69, 65},   // (82) R ,key(69) => A, Ascii Val = 65
		{119, 100, 97}, // w, key(100) => a, Ascii Val = 97
	}

	for _, val := range testCases {
		actualRes := decipherFunc(val.Input, val.CipherKey)
		if val.DecodedRes != actualRes {
			t.Fail()
			t.Errorf("TestDecipherFunc Arguments %v and key %v : FAILED, expected %v but got %v", val.Input, val.CipherKey, val.DecodedRes, actualRes)
		} else {
			t.Logf("TestCipherFunc Arguments %v and key %v : PASSED, expected %v and got %v", val.Input, val.CipherKey, val.DecodedRes, actualRes)
		}
	}

}