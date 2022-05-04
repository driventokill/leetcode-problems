package reversenodesinkgroups

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{newList(1, 2, 3, 4, 5), 2}, newList(2, 1, 4, 3, 5)},
		{"2", args{newList(1, 2, 3, 4, 5), 3}, newList(3, 2, 1, 4, 5)},
		{"3", args{newList(1, 2, 3, 4, 5), 1}, newList(1, 2, 3, 4, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
