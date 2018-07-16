package shuffleAnArray

import (
	"testing"
)

func TestSolution_Shuffle(t *testing.T) {
	tests := []struct {
		name string
		this Solution
		want []int
	}{
		{
			name: "test",
			this: Constructor([]int{2, 1, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.this.Shuffle(), tt.this.Reset(), tt.this.Shuffle(), tt.this.Shuffle())
		})
	}
}
