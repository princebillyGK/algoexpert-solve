package main

import "testing"

type Args struct {
	a   []int64
	sub []int64
}

var tests = []struct {
	name string
	args Args
	want bool
}{
	{
		name: "Test Case 1",
		args: Args{
			a:   []int64{5, 1, 22, 25, 6, -1, 10},
			sub: []int64{1, 6, -1, 10},
		},
		want: true,
	},
	{
		name: "Test Case 2",
		args: Args{
			a:   []int64{5, 1, 22, 25, 6, -1, 10},
			sub: []int64{1, 10, -1},
		},
		want: false,
	},
	{
		name: "Test Case 1",
		args: Args{
			a:   []int64{5, 1, 22, 25, 6, -1, 10},
			sub: []int64{1, 6},
		},
		want: true,
	},
}

func BenchmarkSolveV1(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(t *testing.B) {
			if got := SolveV1(tt.args.a, tt.args.sub); got != tt.want {
				t.Errorf("SolveV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
