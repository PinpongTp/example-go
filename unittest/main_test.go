package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test case 1", args: args{0}, want: 0},
		{name: "test case 2", args: args{1}, want: 1},
		{name: "test case 3", args: args{2}, want: 1},
		{name: "test case 4", args: args{3}, want: 2},
		{name: "test case 5", args: args{4}, want: 3},
		{name: "test case 6", args: args{5}, want: 5},
		{name: "test case 7", args: args{6}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
