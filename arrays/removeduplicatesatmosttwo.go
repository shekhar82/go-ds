package arrays

// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

func removeDuplicates(nums []int) int {

	idxToReplace := 1
	currentCountForSameValue := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			currentCountForSameValue++
		} else {
			currentCountForSameValue = 1
		}

		if currentCountForSameValue <= 2 {
			nums[idxToReplace] = nums[i]
			idxToReplace++
		}
	}
	return idxToReplace
}
