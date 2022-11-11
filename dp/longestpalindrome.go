package dp

func LongestPalindrome(input string) string {
	if len(input) == 0 || len(input) == 1 {
		return input
	} else if len(input) == 2 {
		if input[0] == input[1] {
			return input
		}
		return input[0:1]
	} else {
		longestPalindromeStart, maxLen := 0, 1

		dp := make([][]bool, 0)

		for i := 0; i < len(input); i++ {
			row := make([]bool, len(input))
			for i := 0; i < len(input); i++ {
				row[i] = false
			}
			dp = append(dp, row)
		}

		for i := 0; i < len(input); i++ {
			dp[i][i] = true
		}

		for i := 0; i < len(input)-1; i++ {
			if input[i] == input[i+1] {
				dp[i][i+1] = true
			}
		}

		for i := len(input) - 1; i >= 0; i-- {
			for j := i; j < len(input); j++ {
				dp[i][j] = (input[i] == input[j]) && (((j - i) < 3) || dp[i+1][j-1])

				if dp[i][j] && ((j - i + 1) > maxLen) {
					maxLen = j - i + 1
					longestPalindromeStart = i
				}
			}
		}

		return input[longestPalindromeStart : longestPalindromeStart+maxLen]
	}
}
