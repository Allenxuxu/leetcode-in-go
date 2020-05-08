package algorithm

import "sort"

func permuteUnique(nums []int) [][]int {
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
