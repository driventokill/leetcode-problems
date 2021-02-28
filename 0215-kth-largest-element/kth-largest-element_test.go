// https://leetcode.com/problems/kth-largest-element-in-an-array/

package kthlargestelement

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"2", args{[]int{3, 2, 3, 1, 2, 3, 5, 5, 6}, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
