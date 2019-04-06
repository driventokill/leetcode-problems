// https://leetcode.com/problems/valid-number/

package validnumber

import "testing"

func Test_isNumber(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"empty", " ", false},
		{"leading space", " 1", true},
		{"space", ". 1", false},
		{"tail space", ".1 ", true},
		{"tail space", "1 ", true},
		{"int", "0", true},
		{"float", "0.1", true},
		{"float", "1.", true},
		{"float", ".1", true},
		{"float", "..1", false},
		{"float", ". 1", false},
		{"chars", "abc", false},
		{"chars", "1 e", false},
		{"exponent", "1e3", true},
		{"exponent", "1e", false},
		{"exponent", "99e2.5", false},
		{"exponent", "9.9e25", true},
		{"exponent", "9.9e-25", true},
		{"exponent", "9.9e25e3", false},
		{"exponent", "e9", false},
		{"symbol", "-1", true},
		{"symbol", "--1", false},
		{"symbol", "-+1", false},
		{"symbol", "6+1", false},
		{"symbol", ".+1", false},
		{"symbol", "+.1", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
