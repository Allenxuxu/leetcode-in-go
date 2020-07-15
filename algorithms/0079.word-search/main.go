package algorithm

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}

	mark := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		mark[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, board, mark, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(i, j int, board [][]byte, mark [][]bool, word string, index int) bool {
	if index == len(word)-1 {
		return board[i][j] == word[index]
	}

	if board[i][j] == word[index] {
		mark[i][j] = true

		if i+1 < len(board) && !mark[i+1][j] {
			if dfs(i+1, j, board, mark, word, index+1) {
				return true
			}
		}
		if i-1 >= 0 && !mark[i-1][j] {
			if dfs(i-1, j, board, mark, word, index+1) {
				return true
			}
		}
		if j+1 < len(board[0]) && !mark[i][j+1] {
			if dfs(i, j+1, board, mark, word, index+1) {
				return true
			}
		}
		if j-1 >= 0 && !mark[i][j-1] {
			if dfs(i, j-1, board, mark, word, index+1) {
				return true
			}
		}

		mark[i][j] = false
	}

	return false
}
