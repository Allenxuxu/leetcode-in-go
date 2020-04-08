package algorithm

// 时间复杂度 O(1)
// 空间复杂度 O(1)
// row [9][10]bool 存储每一行(共 9 行)，出现的数字情况（0-9 共10种数字）
// col [9][10]bool 存储每一列(共 9 列)，出现的数字情况（0-9 共10种数字）
// mb  [9][10]bool 存储每一个小矩阵(共 9 个)，出现的数字情况（0-9 共10种数字），(i/3)*3 + j/3 可以计算出当前是第几个小矩阵
// 逐行逐个遍历，如果 board[i][j] 是数字，则判断 此数字是否在 row， col， mb 中出现，如果没出现过，则设置 true
func isValidSudoku(board [][]byte) bool {
	var row [9][10]bool
	var col [9][10]bool
	var mb [9][10]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				number := board[i][j] - '0'
				if row[i][number] {
					return false
				} else {
					row[i][number] = true
				}

				if col[j][number] {
					return false
				} else {
					col[j][number] = true
				}

				index := (i/3)*3 + j/3
				if mb[index][number] {
					return false
				} else {
					mb[index][number] = true
				}
			}
		}
	}

	return true
}
