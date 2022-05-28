package graph

import "testing"

var DefaultGraph = G{
	5, [][]int{
		{0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 1, 1, 0},
	},
}

func TestG_Bfs(t *testing.T) {
	type args struct {
		start int
	}
	tests := []struct {
		name   string
		fields G
		args   args
	}{
		{"1", DefaultGraph, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := G{
				P: tt.fields.P,
				V: tt.fields.V,
			}
			g.Bfs(tt.args.start)
		})
	}
}

func TestG_DfsRecur(t *testing.T) {
	type args struct {
		start int
	}
	tests := []struct {
		name   string
		fields G
		args   args
	}{
		{"1", DefaultGraph, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := G{
				P: tt.fields.P,
				V: tt.fields.V,
			}
			g.DfsRecur(tt.args.start)
		})
	}
}

func TestG_Dfs(t *testing.T) {
	type args struct {
		start int
	}
	tests := []struct {
		name   string
		fields G
		args   args
	}{
		{"1", DefaultGraph, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := G{
				P: tt.fields.P,
				V: tt.fields.V,
			}
			g.Dfs(tt.args.start)
		})
	}
}
