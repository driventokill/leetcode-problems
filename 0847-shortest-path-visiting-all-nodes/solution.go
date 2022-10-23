package shortestpathvisitingallnodes

// struct in queue
// node: node number
// mask: the path to u
// dist: how many step is took before
type NodePath struct {
	node, mask, dist int16
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	if n == 0 {
		return 0
	}

	goal := int16(1<<n) - 1
	var queue = make([]NodePath, 0)

	// [toNode][path] visited
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		// all possible paths to current node
		visited[i] = make([]bool, 1<<n)
		// path of self to self (circle) is zero
		visited[i][1<<i] = true
		queue = append(queue, NodePath{int16(i), 1 << i, 0})
	}

	// BFS
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// shortest path founded
		if node.mask == goal {
			return int(node.dist)
		}

		// search neighbors
		for _, to := range graph[node.node] {
			// add 'to' is visited to origin path
			tempMask := node.mask | 1<<int16(to)
			if !visited[to][tempMask] {
				queue = append(queue, NodePath{int16(to), tempMask, node.dist + 1})
				visited[to][tempMask] = true
			}
		}
	}

	return 0
}
