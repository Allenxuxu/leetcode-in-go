package algorithm

func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = i
	}

	for i := 1; i < len(s); i++ {
		if isPalindrome(s, 0, i) {
			dp[i] = 0
		} else {
			for j := 0; j < i; j++ {
				if isPalindrome(s, j+1, i) {
					dp[i] = min(dp[i], dp[j]+1)
				}
			}
		}
	}

	return dp[len(s)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
