package algorithm

func maxProfit(prices []int) int {
	var ret int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			ret += prices[i+1] - prices[i]
		}
	}
	return ret
}

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, len(prices))
	dp[0] = 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			dp[i] = dp[i-1] + prices[i] - prices[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(prices)-1]
}
