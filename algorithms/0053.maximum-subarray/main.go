package algorithm

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}

	ret := dp[0]
	for i := 1; i < len(nums); i++ {
		ret = max(dp[i], ret)
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	pre, ret := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		pre = max(nums[i], nums[i]+pre)
		ret = max(ret, pre)
	}

	return ret
}
