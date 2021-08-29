package arrays

//Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.

func RemoveDuplicatesInTwoPass(nums []int) int {

	MinInt := -101
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	current := 0
	// 1st loop to replace every duplicates with zero
	for i := 1; i < len(nums); {
		if nums[i] == nums[current] {
			nums[i] = MinInt
			i++
		} else {
			current = i
			i++
		}
	}

	//2nd loop to place position in right place
	current = 1

	for i := 1; i < len(nums); {
		if nums[i] == MinInt {
			i++
		} else {
			nums[current] = nums[i]
			current++
			i++
		}
	}

	return current

}

func RemoveDuplicatesInSinglePass(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	// Variable that would track elments correct position
	pos := 0

	//always point to current non duplicated values
	currentNonDuplicated := 0

	for i := 1; i < len(nums); {
		if nums[i] == nums[currentNonDuplicated] {
			i++
		} else {
			pos = pos + 1
			nums[pos] = nums[i]
			currentNonDuplicated = i
			i++
		}
	}

	return pos + 1
}
