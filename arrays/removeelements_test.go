package arrays

import (
	"sort"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	expectations := []int{0, 0, 1, 3, 4}

	k := RemoveElement(nums, val)

	if k != len(expectations) {
		t.Errorf("Actual length %d, expected was %d", k, len(expectations))
	}

	sort.Ints(nums[0:k])

	for i := 0; i < k; i++ {
		if nums[i] != expectations[i] {
			t.Errorf("expected %d, but was %d", expectations[i], nums[i])
		}
	}
}
