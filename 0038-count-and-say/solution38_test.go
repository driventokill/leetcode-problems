package countandsay

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "n=1",
			input:    1,
			expected: "1",
		},
		{
			name:     "n=2",
			input:    2,
			expected: "11",
		},
		{
			name:     "n=3",
			input:    3,
			expected: "21",
		},
		{
			name:     "n=4",
			input:    4,
			expected: "1211",
		},
		{
			name:     "n=5",
			input:    5,
			expected: "111221",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countAndSay(tt.input)
			if result != tt.expected {
				t.Errorf("countAndSay(%d) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}
