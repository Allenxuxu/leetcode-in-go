package algorithm

func rotate(m [][]int) {
	// n * n
	n := len(m)
	//
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			// [x][y] --> [n - 1 - y][x]
			temp := m[i][j]
			// 左边的行 等于 右边的列
			m[i][j] = m[n-j-1][i]
			m[n-j-1][i] = m[n-i-1][n-j-1]
			m[n-i-1][n-j-1] = m[j][n-i-1]
			m[j][n-i-1] = temp
		}
	}
}
