package _00_number_of_islands

var dir = [][]int{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func inRange(x, y, n, m int) bool {
	return x >= 0 && y >= 0 && x < n && y < m
}

func dfs(grid [][]byte, visited [][]bool, x, y, n, m int) {
	if !inRange(x, y, n, m) || grid[x][y] == '0' || visited[x][y] == true {
		return
	}

	visited[x][y] = true
	for i := 0; i < len(dir); i++ {
		nX := x + dir[i][0]
		nY := y + dir[i][1]

		dfs(grid, visited, nX, nY, n, m)
	}
}

func numIslands(grid [][]byte) int {
	ans := 0
	n := len(grid)
	m := len(grid[0])

	visited := make([][]bool, 0, n)
	for i := 0; i < n; i++ {
		visited = append(visited, make([]bool, m))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == false && grid[i][j] == '1' {
				ans++
				dfs(grid, visited, i, j, n, m)
			}
		}
	}

	return ans
}
