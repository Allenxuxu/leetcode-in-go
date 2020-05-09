package algorithm

func subsets(nums []int) [][]int {
	var result [][]int
	var ans []int
	backtrack(0, nums, &ans, &result)
	return result
}

func backtrack(index int, nums []int, ans *[]int, result *[][]int) {
	tmp := make([]int, len(*ans))
	copy(tmp, *ans)
	*result = append(*result, tmp)

	for i := index; i < len(nums); i++ {
		*ans = append(*ans, nums[i])
		backtrack(i+1, nums, ans, result)
		*ans = (*ans)[:len(*ans)-1]
	}
}
