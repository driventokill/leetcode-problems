package permutations

import (
	"reflect"
	"sort"
	"testing"

	"github.com/driventokill/leetcode-problems/util"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"1", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"2", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{"3", []int{1}, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			sort.Sort(util.IntSlices(tt.want))
			sort.Sort(util.IntSlices(got))

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
