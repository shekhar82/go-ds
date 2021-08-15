package arrays

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	input1 := []int{-4, -1, 0, 3, 10}
	input2 := []int{-7, -3, 2, 3, 11}

	expectedOutput1 := []int{0, 1, 9, 16, 100}
	expectedOutput2 := []int{4, 9, 9, 49, 121}

	actualOutput1 := SortedSquares(input1)
	actualOutput2 := SortedSquares(input2)

	if !reflect.DeepEqual(expectedOutput1, actualOutput1) {
		t.Errorf("expectedOutput1 didn't match with actualOutput1")
	}

	if !reflect.DeepEqual(expectedOutput2, actualOutput2) {
		t.Errorf("expectedOutput2 didn't match with actualOutput2")
	}

}
