// https://leetcode.com/problems/integer-to-roman/
package integertoroman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"3", args{3}, "III"},
		{"4", args{4}, "IV"},
		{"9", args{9}, "IX"},
		{"58", args{58}, "LVIII"},
		{"1994", args{1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
