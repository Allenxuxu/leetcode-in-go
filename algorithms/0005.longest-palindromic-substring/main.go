package algorithm

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var start, end int
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] {
				if j-i > end-start {
					start = i
					end = j
				}
			}
		}
	}

	return s[start : end+1]
}
