package algorithm

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
