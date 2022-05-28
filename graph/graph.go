package graph

import "fmt"

type G struct {
	P int
	V [][]int
}

func (g G) Bfs(start int) {
	visited := make([]int, g.P+1)
	var queue []int
	queue = append(queue, start)
	visited[start] = 1

	for len(queue) != 0 {
		front := queue[0]
		fmt.Println(front)
		queue = queue[1:]

		for i := 1; i <= g.P; i++ {
			if visited[i] == 0 && g.V[front-1][i-1] == 1 {
				visited[i] = 1
				queue = append(queue, i)
			}
		}
	}
}

func (g G) Dfs(start int) {
	visited := make([]int, g.P+1)
	visited[start] = 1
	var stack []int
	stack = append(stack, start)

	for len(stack) > 0 {
		isPush := false

		value := stack[len(stack)-1]
		for i := 1; i <= g.P; i++ {
			if visited[i] == 0 && g.V[value-1][i-1] == 1 {
				visited[i] = 1
				stack = append(stack, i)
				isPush = true
				break
			}
		}

		if !isPush {
			fmt.Println(value)
			stack = stack[:len(stack)-1]
		}
	}
}

func (g G) DfsRecur(start int) {
	visited := make([]int, g.P+1)
	g.DfsRecur0(start, visited)
}

func (g G) DfsRecur0(start int, visited []int) {
	visited[start] = 1

	for i := 1; i <= g.P; i++ {
		if visited[i] == 0 && g.V[start-1][i-1] == 1 {
			g.DfsRecur0(i, visited)
		}
	}
	fmt.Println(start)
}
