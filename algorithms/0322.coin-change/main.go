package algorithm

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 0; i < len(dp); i++ {
		for _, v := range coins {
			if i-v < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-v])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
