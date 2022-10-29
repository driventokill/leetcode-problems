package maxchunkstomakesortesii

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
	tests := []struct {
		name string
		arr []int
		want int
	}{
		{"1", []int{5,4,3,2,1}, 1},
		{"2", []int{2,1,3,4,4}, 4},
		{"3", []int{}, 0},
		{"4", []int{5,3,2,7}, 2},
		{"5", []int{4,2,2,1,1}, 1},
		{"6", []int{1,1,0,0,1}, 2},
		{"7", []int{10,2,3,5,1,6}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunksToSorted(tt.arr); got != tt.want {
				t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
