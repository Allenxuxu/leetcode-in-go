package algorithm

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		if dp[i] > 0 {
			step := nums[i]
			for s := 1; s <= step && i+s < len(nums); s++ {
				if dp[i+s] == 0 {
					dp[i+s] = dp[i] + 1
				}
			}
		}
	}

	return dp[len(nums)-1] - 1
}
