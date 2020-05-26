package algorithm

import "sort"

func permuteUnique(nums []int) [][]int {
	// 排序让重复的数字相邻
	sort.Ints(nums)

	var ret [][]int
	used := make([]bool, len(nums))
	backtrack(nums, &[]int{}, &used, &ret)
	return ret
}

func backtrack(nums []int, current *[]int, used *[]bool, result *[][]int) {
	if len(*current) == len(nums) {
		tmp := make([]int, len(*current))
		copy(tmp, *current)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 剪枝条件：i > 0 是为了保证 nums[i - 1] 有意义
		// 写 !used[i - 1] 是因为 nums[i - 1] 在深度优先遍历的过程中刚刚被撤销选择
		if i != 0 && nums[i-1] == nums[i] && !(*used)[i-1] {
			continue
		}
		if !(*used)[i] {
			*current = append(*current, nums[i])
			(*used)[i] = true
			backtrack(nums, current, used, result)
			(*used)[i] = false
			*current = (*current)[:len(*current)-1]
		}
	}
}
