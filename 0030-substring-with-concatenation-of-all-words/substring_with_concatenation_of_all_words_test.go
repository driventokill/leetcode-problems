package substring_with_concatenation_of_all_words

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"abcefgcsfg", []string{"abc"}}, []int{0}},
		{"2", args{"barfoothefoobarman", []string{"foo", "bar"}}, []int{0, 9}},
		{"3", args{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}}, []int{}},
		{"4", args{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}}, []int{6, 9, 12}},
		{"5", args{"wordgoodgoodgoodbestword", []string{"word","good","best","good"}}, []int{8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
