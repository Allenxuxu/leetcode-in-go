package algorithm

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		if dp[i] == 1 {
			step := nums[i]
			for s := 1; s <= step && i+s < len(nums); s++ {
				dp[i+s] = 1
			}
		}
	}

	return dp[len(nums)-1] == 1
}
