package main

import (
	"reflect"
	"testing"
)

func TestRecoverFromPreorder(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *TreeNode
	}{
		{
			name:  "单节点树",
			input: "1",
			expected: &TreeNode{
				Val: 1,
			},
		},
		{
			name:  "有左子树的树",
			input: "1-2",
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			name:  "完整的示例",
			input: "1-2--3--4-5--6--7",
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := recoverFromPreorder(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("recoverFromPreorder() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 辅助函数：用于比较两个树是否相等
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
