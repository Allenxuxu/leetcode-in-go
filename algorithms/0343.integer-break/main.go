package algorithm

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 1

	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), dp[i-j]*j))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
