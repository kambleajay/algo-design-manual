package stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}
	for _, test := range tests {
		if got := climbStairs(test.n); got != test.want {
			t.Errorf("climbStairs(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
