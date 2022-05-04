package validsudoku

func isValidSudoku(board [][]byte) bool {
	hSet := make([][]bool, 9)
	vSet := make([][]bool, 9)
	cSet := make([][]bool, 9)

	// init sets
	for i := 0; i < 9; i++ {
		hSet[i] = make([]bool, 9)
		vSet[i] = make([]bool, 9)
		cSet[i] = make([]bool, 9)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == byte('.') {
				continue
			}

			value := board[i][j] - byte('1')

			if hSet[i][value] {
				return false
			}
			hSet[i][value] = true

			if vSet[j][value] {
				return false
			}
			vSet[j][value] = true

			rr := (i/3)*3 + j/3
			if cSet[rr][value] {
				return false
			}
			cSet[rr][value] = true
		}
	}

	return true
}
