package algorithm

func rotate(m [][]int) {
	// n * n
	n := len(m)
	// 从外往内逐层缩小，终止条件 n/2 。 例如，4 * 4 --> 3 * 3 --> 2 * 2
	for i := 0; i < n/2; i++ {
		// 每一层需要的旋转的个数，每层元素个数 - 1 。 n - i 为当前缩小后的矩阵 层数。
		for j := i; j < n-i-1; j++ {
			// 找出索引变换的规律 [x][y] --> [n - 1 - y][x]
			// 旋转 4 个元素
			temp := m[i][j]
			// 左边的行 等于 右边的列
			m[i][j] = m[n-j-1][i]
			m[n-j-1][i] = m[n-i-1][n-j-1]
			m[n-i-1][n-j-1] = m[j][n-i-1]
			m[j][n-i-1] = temp
		}
	}
}
