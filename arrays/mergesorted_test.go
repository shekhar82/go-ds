package arrays

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	num1 := []int{0}
	num2 := []int{1}

	expected := []int{1}

	Merge(num1, 0, num2, 1)

	if !reflect.DeepEqual(expected, num1) {
		t.Errorf("Test case failed")
	}

}
