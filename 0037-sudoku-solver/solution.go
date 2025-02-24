package sudokusolver

func solveSudoku(board [][]byte) {
	backtrack(board, 0, 0)
}

func backtrack(board [][]byte, r, c int) bool {
	var m, n int = 9, 9
	if c == n {
		// reached the last column, move to the next row
		return backtrack(board, r+1, 0)
	}

	if r == m {
		return true
	}

	// if current position is already setted, no need to set
	if board[r][c] != '.' {
		return backtrack(board, r, c+1)
	}

	for i := 1; i <= 9; i++ {
		// invalid number for current position
		if !isValid(board, r, c, byte('0'+i)) {
			continue
		}

		board[r][c] = byte('0' + i)
		if backtrack(board, r, c+1) {
			return true
		}
		board[r][c] = '.'
	}
	return false
}

// isValid returns if `n` can be put into row `r` column `c`
func isValid(board [][]byte, r, c int, n byte) bool {
	for i := 0; i < 9; i++ {
		if board[r][i] == n {
			return false
		}
		if board[i][c] == n {
			return false
		}
		if board[r/3*3+i/3][c/3*3+i%3] == n {
			return false
		}
	}

	return true
}
