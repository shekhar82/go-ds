package arrays

//Given a binary array nums, return the maximum number of consecutive 1's in the array.

//Example 1
//Input: nums = [1,1,0,1,1,1]
//Output: 3
//Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.

//Example 1
//Input: nums = [1,0,1,1,0,1]
//Output: 2

func FindMaxConsecutiveOnes(nums []int) int {
	maxOccurrence := 0

	currentOneSprint := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if currentOneSprint > maxOccurrence {
				maxOccurrence = currentOneSprint
			}
			currentOneSprint = 0
		} else if nums[i] == 1 {
			currentOneSprint += 1
		}
	}

	if currentOneSprint > maxOccurrence {
		maxOccurrence = currentOneSprint
	}
	return maxOccurrence
}
