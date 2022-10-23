package count_bits

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{0}, []int{0}},
		{"1", args{2}, []int{0,1,1}},
		{"2", args{5}, []int{0,1,1,2,1,2}},
		{"8", args{8}, []int{0,1,1,2,1,2,2,3,1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
