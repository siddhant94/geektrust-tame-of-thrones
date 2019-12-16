package messages

import "testing"

func TestComparator(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 1}
	m2 := map[int]int{1: 2, 3: 1} //  True Case:Both maps are same in keys && value of m1 <= val of m2
	m3 := map[int]int{4: 1, 65: 1}
	m4 := map[int]int{5: 1, 65: 1} // False Case: Different Keys
	m5 := map[int]int{87: 3, 66: 4}
	m6 := map[int]int{66: 4, 87: 2} // False Case: same keys with val m1 > val m2
	m7 := map[int]int{1: 1, 2: 1}
	m8 := map[int]int{2: 1, 1: 1} // True Case: same map i.e. nothing is different

	dataItems := []TestComparatorItem{
		{[2]map[int]int{m1, m2}, true},
		{[2]map[int]int{m3, m4}, false},
		{[2]map[int]int{m5, m6}, false},
		{[2]map[int]int{m7, m8}, true},
	}

	for _, val := range dataItems {
		actualResult := comparator(val.inputs[0], val.inputs[1])
		if actualResult != val.result {
			t.Fail()
			t.Errorf("TestComparator Arguments %v : FAILED, expected %v but got %v", val.inputs, val.result, actualResult)
		} else {
			t.Logf("TestComparator Arguments %v : PASSED, expected %v and got %v", val.inputs, val.result, actualResult)
		}
	}
}
