package nqueens

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)

	// init board with '.'
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	backtrack(&res, board, 0)

	return res
}

func backtrack(res *[][]string, board [][]rune, row int) {
	if row == len(board) {
		answer := make([]string, len(board))
		for i, v := range board {
			answer[i] = string(v)
		}
		*res = append(*res, answer)
		return
	}

	for col := range board[row] {
		if !isValid(board, row, col) {
			continue
		}

		// 1. make a decision
		board[row][col] = 'Q'
		// 2. next step based on current dicision
		backtrack(res, board, row+1)
		// 3. revert and go next
		board[row][col] = '.'
	}
}

func isValid(board [][]rune, row, col int) bool {
	// 1. check conflicts by column
	for i := 0; i < len(board); i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 2. check conflicts by row(no need, because we decide by row)
	// 3. check conflicts by corner
	// 3.1 check right top
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	//3.2 check left top
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}
