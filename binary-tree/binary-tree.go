package binarytree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const EMPTY_NODE = math.MinInt

// NewTree constructs binary tree from value slice.
// [0,1,2,3,4,5,6]
//
//	     0
//	   /   \
//	  1     2
//	 / \   / \
//	3   4 5   6
func NewTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{nums[0], nil, nil}
	var queue []*TreeNode = []*TreeNode{root}
	var values = nums[1:]

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if len(values) >= 2 {
			left := &TreeNode{values[0], nil, nil}
			right := &TreeNode{values[1], nil, nil}
			node.Left = left
			node.Right = right
			queue = append(queue, left, right)
			values = values[2:]
		} else if len(values) == 1 {
			left := &TreeNode{values[0], nil, nil}
			node.Left = left
			queue = append(queue, left)
			values = values[1:]
		} else {
			break
		}
	}

	return root
}
