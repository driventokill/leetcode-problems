package combinationsum

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{2, 3, 6, 7}, 7}, [][]int{{2, 2, 3}, {7}}},
		{"2", args{[]int{2, 3, 5}, 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"3", args{[]int{2}, 1}, [][]int{}},
		{"4", args{[]int{2, 7, 6, 3, 5, 1}, 9}, [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 2}, {1, 1, 1, 1, 1, 1, 3}, {1, 1, 1, 1, 1, 2, 2}, {1, 1, 1, 1, 2, 3},
			{1, 1, 1, 1, 5}, {1, 1, 1, 2, 2, 2}, {1, 1, 1, 3, 3}, {1, 1, 1, 6}, {1, 1, 2, 2, 3},
			{1, 1, 2, 5}, {1, 1, 7}, {1, 2, 2, 2, 2}, {1, 2, 3, 3}, {1, 2, 6}, {1, 3, 5}, {2, 2, 2, 3}, {2, 2, 5}, {2, 7}, {3, 3, 3}, {3, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
