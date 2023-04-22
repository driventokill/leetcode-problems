package allpossiblefullybinarytrees

import (
	"strings"
	"testing"
)

func Test_allPossibleFBT(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPossibleFBT(tt.args.n); !TreeEquals(got, tt.want) {
				t.Errorf("allPossibleFBT() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func NewTree(vals []*int) *TreeNode {
// 	if len(vals) == 0 {
// 		return nil
// 	}

// 	root := &TreeNode{
// 		Val: *vals[0],
// 		Left: nil,
// 		Right: nil,
// 	}

// 	for i := 1; i < len(vals); i++ {

// 	}
// }

func TreeEquals(t1 []*TreeNode, t2 []*TreeNode) bool {
	m1 := make(map[string]bool)
	m2 := make(map[string]bool)

	for _, v := range t1 {
		m1[Tree2String(v)] = true
	}

	for _, v := range t2 {
		str := Tree2String(v)
		if _, exist := m1[str]; !exist {
			return false
		}
		m2[str] = true
	}

	if len(m1) != len(m2) {
		return false
	}

	return true
}

func Tree2String(root *TreeNode) string {
	if root == nil {
		return " "
	}

	return strings.Join([]string{Tree2String(root.Left), string(root.Val), Tree2String(root.Right)}, "")
}
