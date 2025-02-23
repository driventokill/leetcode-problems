package maximumscorefromremovingsubstrings

import "testing"

func TestMaximumGain(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		x        int
		y        int
		expected int
	}{
		{
			name:     "case1: ab > ba",
			s:        "cdbcbbaaabab",
			x:        4,
			y:        2,
			expected: 12,
		},
		{
			name:     "case2: ba > ab",
			s:        "aabbaaxybbaabb",
			x:        3,
			y:        4,
			expected: 16,
		},
		{
			name:     "case3: empty string",
			s:        "",
			x:        5,
			y:        3,
			expected: 0,
		},
		{
			name:     "case4: no valid substrings",
			s:        "xyz",
			x:        2,
			y:        3,
			expected: 0,
		},
		{
			name:     "case5: single type of substring",
			s:        "ababab",
			x:        4,
			y:        2,
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumGain(tt.s, tt.x, tt.y)
			if got != tt.expected {
				t.Errorf("maximumGain(%q, %d, %d) = %d, want %d",
					tt.s, tt.x, tt.y, got, tt.expected)
			}
		})
	}
}
