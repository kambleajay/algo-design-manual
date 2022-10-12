package recursion

import (
	"testing"
)

func TestNumSquares(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{12, 3},
		{13, 2},
		{25, 1},
		{27, 3},
	}
	for _, test := range tests {
		if got := numSquares(test.n); got != test.want {
			t.Errorf("numSquares(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
