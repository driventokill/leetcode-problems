package binarytree

import (
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{"1", []int{1, 2, 3}, &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTree(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
