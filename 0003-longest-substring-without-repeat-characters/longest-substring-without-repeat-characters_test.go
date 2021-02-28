package longestsubstringwithoutrepeatcharacters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"abcabcbb"}, 3},
		{"Example 2", args{"bbbbb"}, 1},
		{"Example 3", args{"pwwkew"}, 3},
		{"Example 4", args{""}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring1(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring1() = %v, want %v", got, tt.want)
			}
		})
	}
}
