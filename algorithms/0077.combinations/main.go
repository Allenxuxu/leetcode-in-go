package algorithm

func combine(n int, k int) [][]int {
	var ret [][]int
	backtrack(n, k, 1, &[]int{}, &ret)
	return ret
}

func backtrack(n, k, index int, current *[]int, result *[][]int) {
	if k == len(*current) {
		tmp := make([]int, len(*current))
		copy(tmp, *current)
		*result = append(*result, tmp)
		return
	}

	for i := index; i <= n; i++ {
		*current = append(*current, i)
		backtrack(n, k, i+1, current, result)
		*current = (*current)[:len(*current)-1]
	}
}
