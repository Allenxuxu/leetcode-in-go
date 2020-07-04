package algorithm

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var min, max = prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if max < prices[i]-min {
				max = prices[i] - min
			}
		}
	}
	return max
}
