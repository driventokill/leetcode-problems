// https://leetcode.com/problems/roman-to-integer/

package romantointeger

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{""}, 0},
		{"3", args{"III"}, 3},
		{"4", args{"IV"}, 4},
		{"9", args{"IX"}, 9},
		{"13", args{"XIII"}, 13},
		{"58", args{"LVIII"}, 58},
		{"1994", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
