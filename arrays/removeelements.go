package arrays

//Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The relative order of the elements may be changed.

func RemoveElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] != val {
			return 1
		} else {
			return 0
		}
	}

	i, j := 0, len(nums)-1

	for i < j {
		if nums[i] == val && nums[j] != val {
			nums[i] = nums[j]
			nums[j] = val
			i++
			j--
		} else if nums[i] != val {
			i++
		} else if nums[j] == val {
			j--
		}
	}

	if i == 0 && nums[i] == val {
		return 0
	} else if nums[i] != val {
		return i + 1
	} else {
		return i
	}
}
