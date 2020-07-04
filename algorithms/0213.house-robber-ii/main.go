package algorithm

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(robR(nums[:len(nums)-1]), robR(nums[1:]))
}
func robR(nums []int) int {
	f, s := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp := max(s, f+nums[i])
		f = s
		s = tmp
	}

	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
