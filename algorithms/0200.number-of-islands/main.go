package algorithm

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var ret int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				ret++
			}
		}
	}

	return ret
}

func dfs(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	// 上
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
	// 下
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	// 左
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
	// 右
	if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
}
