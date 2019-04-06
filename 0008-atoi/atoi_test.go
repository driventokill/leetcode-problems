package atoi

import "testing"

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{"symbol", "+-1", 0},
		{"symbol", "-1", -1},
		{"with words", "4193 the words", 4193},
		{"upper flow", "2873498734182581725", 1<<31 - 1},
		{"lower flow", "-238482375873285712857", -1 << 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
