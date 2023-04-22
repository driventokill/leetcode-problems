package mergeksortedlists

import (
	"reflect"
	"testing"
)

import . "github.com/driventokill/leetcode-problems/util"

func Test_mergeKListsByHeap(t *testing.T) {
	type args struct {
		lists [][]int
	}
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{"1", [][]int{{2, 3, 7}, {1, 2, 5}, {9, 10}}, []int{1, 2, 2, 3, 5, 7, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := List2Slice(mergeKListsByHeap(Ints2Lists(tt.lists))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKListsByHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
