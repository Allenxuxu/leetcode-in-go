package algorithm

func permute(nums []int) [][]int {
	var ret [][]int

	backtrack(nums, &[]int{}, &ret)
	return ret
}

func backtrack(nums []int, current *[]int, result *[][]int) {
	if len(*current) == len(nums) {
		tmp := make([]int, len(*current))
		copy(tmp, *current)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !has(*current, nums[i]) {
			*current = append(*current, nums[i])
			backtrack(nums, current, result)
			*current = (*current)[:len(*current)-1]
		}
	}
}

func has(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}

	return false
}
