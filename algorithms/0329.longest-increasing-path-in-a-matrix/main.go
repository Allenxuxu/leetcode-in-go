package algorithm

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		m[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			m[i][j] = 1
		}
	}

	var ret int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ret = max(ret, dfs(i, j, matrix, m))
		}
	}

	return ret
}

func dfs(i, j int, matrix [][]int, m [][]int) int {
	if m[i][j] != 1 {
		return m[i][j]
	}

	if j-1 >= 0 && matrix[i][j-1] > matrix[i][j] { // 左
		m[i][j] = max(m[i][j], 1+dfs(i, j-1, matrix, m))
	}
	if j+1 < len(matrix[i]) && matrix[i][j+1] > matrix[i][j] { // 右
		m[i][j] = max(m[i][j], 1+dfs(i, j+1, matrix, m))
	}
	if i-1 >= 0 && matrix[i-1][j] > matrix[i][j] { //上
		m[i][j] = max(m[i][j], 1+dfs(i-1, j, matrix, m))
	}
	if i+1 < len(matrix) && matrix[i+1][j] > matrix[i][j] { //下
		m[i][j] = max(m[i][j], 1+dfs(i+1, j, matrix, m))
	}

	return m[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
