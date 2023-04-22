package zigzagconversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]rune, numRows)

	for i := 0; i < numRows; i++ {
		rows[i] = make([]rune, 0)
	}

	row := 0
	down := true
	for _, r := range s {
		rows[row] = append(rows[row], r)
		if down { // going down
			if row == numRows-1 { // reach bottom
				down = !down
				row = row - 1
			} else {
				row = row + 1
			}
		} else { // going up
			if row == 0 { // reach top
				down = !down
				row = row + 1
			} else {
				row = row - 1
			}
		}
	}

	ret := make([]rune, 0)
	for _, row := range rows {
		ret = append(ret, row...)
	}

	return string(ret)
}
