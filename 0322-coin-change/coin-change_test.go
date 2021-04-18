package _322_coin_change

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3", args{[]int{1, 2, 5}, 11}, 3},
		{"-1", args{[]int{2}, 3}, -1},
		{"0", args{[]int{1}, 0}, 0},
		{"1", args{[]int{1}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}

			if got := coinChange1(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange1() = %v, want %v", got, tt.want)
			}
		})
	}
}
