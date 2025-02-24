package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
	if len(traversal) == 0 {
		return nil
	}

	var root *TreeNode

	var depth int = 0
	var path []*TreeNode
	var val int

	for _, c := range traversal {

		// Collect value from characters
		if c != '-' {
			val = val*10 + int(c-'0')
			continue
		}

		// Enter the '-' substring zone

		// Initialize the root node first. Will only execute once.
		if root == nil {
			root = &TreeNode{Val: val}
			path = append(path, root)
			// Reset val after inserting
			val = 0
			depth = 0
			continue
		}

		depth++

		if val > 0 {
			// Handle the value collected from string for current node
			// insert to right place
			parent := path[depth-1]
			n := &TreeNode{Val: val}
			if parent.Left == nil {
				parent.Left = n
			} else {
				parent.Right = n
			}

			path = append(path[:depth], n)

			// Reset status of depth and val
			val = 0
			depth = 0
		}

	}

	if val > 0 {
		if root == nil {
			root = &TreeNode{Val: val}
		} else {
			parent := path[depth]
			n := &TreeNode{Val: val}
			if parent.Left == nil {
				parent.Left = n
			} else {
				parent.Right = n
			}
		}
	}

	return root
}

func recoverFromPreorder1(traversal string) *TreeNode {

	nodes := strings.Split(traversal, "-")
	if len(nodes) == 0 {
		return nil
	}

	fmt.Printf("%v split to : %v\n", traversal, nodes)

	root := appendNode(nil, nodes[0])

	var depthNode []*TreeNode
	var depth int = 1

	depthNode = append(depthNode, root)

	for i := 1; i < len(nodes); i++ {
		if nodes[i] == "" {
			depth = depth + 1
			continue
		}

		currentRoot := depthNode[depth-1]

		newRoot := appendNode(currentRoot, nodes[i])

		if len(depthNode) < depth+1 {
			depthNode = append(depthNode, newRoot)
		} else {
			depthNode[depth] = newRoot
		}
		// reset depth
		depth = 1
	}

	return root
}

func appendNode(root *TreeNode, value string) *TreeNode {
	i, _ := strconv.Atoi(value)
	node := &TreeNode{Val: i}

	if root == nil {
		return node
	}

	if root.Left == nil {
		root.Left = node
		return node
	}

	if root.Right == nil {
		root.Right = node
		return node
	}

	return nil
}

func (n TreeNode) String() string {
	return fmt.Sprintf("{Val: %v, Left: %v, Right: %v}", n.Val, n.Left, n.Right)
}
