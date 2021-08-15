package arrays

//Given an array nums of integers, return how many of them contain an even number of digits.

func FindNumbers(nums []int) int {

	evenDigitCounts := 0

	for _, val := range nums {
		if isEven(getCountOfDigits(val)) {
			evenDigitCounts = evenDigitCounts + 1
		}
	}

	return evenDigitCounts
}

func isEven(num int) bool {
	return num%2 == 0
}

func getCountOfDigits(num int) int {
	count := 0
	val := num
	for val != 0 {
		val = val / 10
		count = count + 1
	}
	return count
}
