package messages

import (
	"reflect"
	"testing"
)

type TestComparatorItem struct {
	inputs [2]map[int]int
	result bool
}

func TestReadInput(t *testing.T) {
	var expected []string
	expected = []string{"Hello, this is a test input", "This is scond line in the file with line-break"}
	inputFilename := "testInput1.txt"
	result := ReadFileInput(inputFilename)
	if !reflect.DeepEqual(expected, result) { // Reflect only used for testing purpose as bad performance
		t.Fail()
		t.Errorf("TestReadInput Arguments %s : FAILED, expected %v but got %v", inputFilename, expected, result)
	} else {
		t.Logf("TestReadInput Arguments %s : PASSED, expected %v and got %v", inputFilename, expected, result)
	}
}

func TestParseInputData(t *testing.T) {
	// [ICE MAMMOTH THVAO FIRE DRAGON JXGMUT AIR ZORRO WATER OCTO VJAVWBZ PUS]
	var testCases = []struct {
		Input  []string
		Result map[string]string
	}{
		{Input: []string{"ICE MAMMOTH THVAO", "FIRE DRAGON JXGMUT", "AIR ZORRO", "WATER OCTO VJAVWBZ PUS"},
			Result: map[string]string{
				"ICE":   "MAMMOTHTHVAO",
				"FIRE":  "DRAGONJXGMUT",
				"AIR":   "ZORRO",
				"WATER": "OCTOVJAVWBZPUS",
			},
		},
	}
	for _, val := range testCases {
		res := ParseInputData(val.Input)
		if reflect.DeepEqual(res, val.Result) {
			t.Fail()
			t.Errorf("TestParseInputData Arguments %v : FAILED, expected %v but got %v", val.Input, val.Result, res)
		} else {
			t.Logf("TestParseInputData Arguments %v : PASSED, expected %v and got %v", val.Input, val.Result, res)
		}
	}

}

func TestProcessMessages(t *testing.T) {
	var testCases = []struct {
		Input    map[string]string
		Expected []string
	}{
		{Input: map[string]string{
			"ICE":   "MAMMOTHTHVAO",
			"FIRE":  "DRAGONJXGMUT",
			"AIR":   "ZORRO",
			"WATER": "OCTOVJAVWBZPUS",
		},
			Expected: []string{"FIRE", "AIR", "WATER"},
		},
	}
	for _, val := range testCases {
		res := ProcessMessages(val.Input)
		if reflect.DeepEqual(res, val.Expected) {
			t.Fail()
			t.Errorf("TestProcessMessages Arguments %v : FAILED, expected %v but got %v", val.Input, val.Expected, res)
		} else {
			t.Logf("TestProcessMessages Arguments %v : PASSED, expected %v and got %v", val.Input, val.Expected, res)
		}
	}

}
