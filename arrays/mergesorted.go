package arrays

func Merge(num1 []int, m int, num2 []int, n int) {
	putZeroAhead(&num1, m, n)

	num1Idx := n
	num2Idx := 0

	pos := 0

	for {
		if num1Idx >= (m+n) || num2Idx >= n {
			break
		}

		if num1[num1Idx] <= num2[num2Idx] {
			num1[pos] = num1[num1Idx]
			pos++
			num1Idx++
		} else {
			num1[pos] = num2[num2Idx]
			pos++
			num2Idx++
		}
	}

	if num2Idx < n {
		for i := num2Idx; i < n; i++ {
			num1[pos] = num2[i]
			pos++
		}
	}

}

func putZeroAhead(p *[]int, m int, n int) {
	pos := m + n - 1

	for i := m - 1; i >= 0; i-- {
		(*p)[pos] = (*p)[i]
		pos--
	}

	for i := pos; i >= 0; i-- {
		(*p)[i] = 0
	}

}
