package numberofwaystobuypencilsandpens

import "testing"

func Test_waysToBuyPensPencils(t *testing.T) {
	type args struct {
		total int
		cost1 int
		cost2 int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{20,10,5}, 9},
		{"2", args{5, 10, 10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToBuyPensPencils(tt.args.total, tt.args.cost1, tt.args.cost2); got != tt.want {
				t.Errorf("waysToBuyPensPencils() = %v, want %v", got, tt.want)
			}
		})
	}
}
