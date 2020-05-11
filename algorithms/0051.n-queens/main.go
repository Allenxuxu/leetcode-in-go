package algorithm

import "bytes"

func solveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}
	var result [][]string
	board := make([][]byte, n)
	for i := 0; i < len(board); i++ {
		tmpRow := bytes.Repeat([]byte{'.'}, len(board))
		board[i] = tmpRow
	}

	backtrack(&board, 0, &result)

	return result
}

func backtrack(board *[][]byte, row int, result *[][]string) {
	if row == len(*board) {
		tmp := make([]string, 0, len(*board))
		for _, v := range *board {
			tmp = append(tmp, string(v))
		}
		*result = append(*result, tmp)
	}

	for col := 0; col < len(*board); col++ {
		if !isValid(board, row, col) {
			continue
		}

		(*board)[row][col] = 'Q'
		backtrack(board, row+1, result)
		(*board)[row][col] = '.'
	}
}

func isValid(board *[][]byte, row, col int) bool {
	// 检查列
	for i := 0; i < len(*board); i++ {
		if (*board)[i][col] == 'Q' {
			return false
		}
	}

	// 检查右上
	for i, j := row-1, col+1; i >= 0 && j < len(*board); {
		if (*board)[i][j] == 'Q' {
			return false
		}

		i--
		j++
	}

	// 检查左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if (*board)[i][j] == 'Q' {
			return false
		}

		i--
		j--
	}

	return true
}
