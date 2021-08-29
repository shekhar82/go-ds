package arrays

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	expectations := []int{1, 2}

	k := RemoveDuplicatesInSinglePass(nums)

	if k != len(expectations) {
		t.Errorf("Actual length %d, expected was %d", k, len(expectations))
	}

	//sort.Ints(nums[0:k])

	for i := 0; i < k; i++ {
		if nums[i] != expectations[i] {
			t.Errorf("expected %d, but was %d", expectations[i], nums[i])
		}
	}
}
