package minimumdepthofbinarytree

import (
	binarytree "github.com/driventokill/leetcode-problems/binary-tree"
)

func minDepth(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*binarytree.TreeNode = []*binarytree.TreeNode{root}

	var depth = 1

	for len(queue) > 0 {
		cnt := len(queue)

		for i := 0; i < cnt; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}

	return depth
}
