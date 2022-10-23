package shortestpathvisitingallnodes

import "testing"

func Test_shortestPathLength(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  int
	}{
		{"1", [][]int{{1, 2, 3}, {0}, {0}, {0}}, 4},
		{"2", [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathLength(tt.graph); got != tt.want {
				t.Errorf("shortestPathLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
