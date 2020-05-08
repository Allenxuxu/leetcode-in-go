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

func permute1(nums []int) [][]int {
	var ret [][]int
	used := make([]bool, len(nums))
	backtrack1(nums, &[]int{}, &used, &ret)
	return ret
}

func backtrack1(nums []int, current *[]int, used *[]bool, result *[][]int) {
	if len(*current) == len(nums) {
		tmp := make([]int, len(*current))
		copy(tmp, *current)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			*current = append(*current, nums[i])
			(*used)[i] = true
			backtrack1(nums, current, used, result)
			(*used)[i] = false
			*current = (*current)[:len(*current)-1]
		}
	}
}
