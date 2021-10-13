package main

import "fmt"

func main() {
	input := []int{-5, -3, -2, -1}
	output := sortedSquares(input)

	// Testing Git commit
	fmt.Println(output)
}

func sortedSquares(nums []int) []int {
	finalArray := make([]int, len(nums))

	idx := 0

	if nums[0] >= 0 {
		for index, _ := range nums {
			nums[index] = nums[index] * nums[index]
		}

		return nums
	} else {
		i := 0
		for ; i < len(nums); i++ {
			if nums[i] >= 0 {
				break
			} else {
				nums[i] = abs(nums[i])
			}
		}

		j := i - 1

		for i < len(nums) && j >= 0 {
			if nums[i] > nums[j] {
				finalArray[idx] = nums[j] * nums[j]
				j--
			} else {
				finalArray[idx] = nums[i] * nums[i]
				i++
			}
			idx++
		}

		if i < len(nums) {
			for ; i < len(nums); i++ {
				finalArray[idx] = nums[i] * nums[i]
				idx++
			}

		}

		if j >= 0 {
			for ; j >= 0; j-- {
				finalArray[idx] = nums[j] * nums[j]
				idx++
			}

		}

		return finalArray
	}

}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
