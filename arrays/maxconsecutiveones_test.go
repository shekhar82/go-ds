package arrays

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	input1 := []int{1, 1, 0, 1, 1, 1}
	input2 := []int{1, 0, 1, 1, 0, 1}

	expectedOutput1 := 3
	expectedOutput2 := 2

	actualOutput1 := FindMaxConsecutiveOnes(input1)
	actualOutput2 := FindMaxConsecutiveOnes(input2)

	if actualOutput1 != expectedOutput1 {
		t.Errorf("FindMaxConsecutiveOnes = %d; want %d", actualOutput1, expectedOutput1)
	}

	if actualOutput2 != expectedOutput2 {
		t.Errorf("FindMaxConsecutiveOnes = %d; want %d", actualOutput2, actualOutput2)
	}

}
