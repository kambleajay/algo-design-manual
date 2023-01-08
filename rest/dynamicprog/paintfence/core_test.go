package paintfence

import "testing"

func TestNumWays(t *testing.T) {
	var tests = []struct {
		n    int
		k    int
		want int
	}{
		{3, 2, 6},
		{1, 1, 1},
		{7, 2, 42},
	}
	for _, test := range tests {
		if got := numWays(test.n, test.k); got != test.want {
			t.Errorf("numWays(%d, %d) = %d, want %d", test.n, test.k, got, test.want)
		}
	}
}
