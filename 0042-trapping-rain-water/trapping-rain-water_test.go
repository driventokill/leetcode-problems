// https://leetcode.com/problems/trapping-rain-water/

package trappingrainwater

import "testing"

func Test_trap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"test", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"test", []int{1, 2, 0, 1, 0, 3, 1}, 5},
		{"test", []int{2, 0, 1, 3, 1, 0, 2}, 6},
		{"test", []int{2, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
