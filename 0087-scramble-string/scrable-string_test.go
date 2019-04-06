package scramblestring

import "testing"

func Test_isScramble(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ab", args{"ab", "ba"}, true},
		{"abc", args{"abc", "bac"}, true},
		{"abc", args{"abc", "abc"}, true},
		{"abc", args{"abc", "bca"}, true},
		{"great", args{"great", "rgeat"}, true},
		{"great", args{"great", "rgtae"}, true},
		{"abcde", args{"abcde", "edcba"}, true},
		{"test", args{"ababcbaccbabbcbca", "bbbbbaaaacccccbba"}, false},
		{"longstring", args{"ababcbaccbabbcbcaabddagdasgadasbasadsgasdg", "abddagdasgadasbasadsgasdgababcbaccbabbcbca"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}
