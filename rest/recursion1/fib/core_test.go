package fib

import "testing"

func TestFib(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}
	for _, test := range tests {
		if got := fib(test.n); got != test.want {
			t.Errorf("fib(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
