package allpossiblefullybinarytrees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}

	if n == 1 {
		return []*TreeNode{new(TreeNode)}
	}

	var nodes []*TreeNode

	for i := 0; i < n-1; i++ {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - i - 1)
		for _, leftNode := range left {
			for _, rightNode := range right {
				root := new(TreeNode)
				root.Left = leftNode
				root.Right = rightNode
				nodes = append(nodes, root)
			}
		}
	}

	return nodes
}
