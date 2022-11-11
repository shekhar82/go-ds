package systemdesign

func nextPermutation(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	i := len(nums) - 2

	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}

	reverse(nums, i+1)
	return nums
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, start int) {
	i := start
	j := len(nums) - 1

	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}
