package minimumdepthofbinarytree

import (
	"testing"

	binarytree "github.com/driventokill/leetcode-problems/binary-tree"
)

func Test_minDepth(t *testing.T) {
	tests := []struct {
		name string
		root *binarytree.TreeNode
		want int
	}{
		{"1", &binarytree.TreeNode{3, &binarytree.TreeNode{9, nil, nil}, &binarytree.TreeNode{20, &binarytree.TreeNode{15, nil, nil}, &binarytree.TreeNode{7, nil, nil}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
