package countspecialquadruplets

import "testing"

func Test_countQuadruplets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{1, 2, 3, 6}}, 1},
		{"3", args{[]int{1, 1, 1, 3, 5}}, 4},
		{"3", args{[]int{2, 16, 9, 27, 3, 39}}, 2},
		{"3", args{[]int{28, 8, 49, 85, 37, 90, 20, 8}}, 1},
		{"3", args{[]int{35, 15, 38, 1, 10, 26}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countQuadruplets(tt.args.nums); got != tt.want {
				t.Errorf("countQuadruplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
