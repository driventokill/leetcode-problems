package minimumfallingpathsum

import "math"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	res := math.MaxInt

	m := newMemo(n)

	for i := 0; i < n; i++ {
		v := m.dp(matrix, n-1, i)
		if v < res {
			res = v
		}
	}

	return res
}

type WithMemo struct {
	memo [][]int
}

func newMemo(size int) *WithMemo {
	m := new(WithMemo)
	memo := make([][]int, size)

	for i := range memo {
		memo[i] = make([]int, size)
		for j := range memo[i] {
			memo[i][j] = math.MaxInt
		}
	}

	m.memo = memo
	return m
}

// dp calculate the mininum path sum from matrix[0][...] to matrix[i][j]
func (m *WithMemo) dp(matrix [][]int, i, j int) int {
	// range validation
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix) {
		return math.MaxInt
	}

	if i == 0 {
		m.memo[i][j] = matrix[i][j]
		return matrix[i][j]
	}

	if m.memo[i][j] != math.MaxInt {
		return m.memo[i][j]
	}

	minimum := matrix[i][j] + min(
		m.dp(matrix, i-1, j-1),
		m.dp(matrix, i-1, j),
		m.dp(matrix, i-1, j+1),
	)
	m.memo[i][j] = minimum

	return minimum
}

func min(i ...int) int {
	v := math.MaxInt

	for _, item := range i {
		if item < v {
			v = item
		}
	}

	return v
}
