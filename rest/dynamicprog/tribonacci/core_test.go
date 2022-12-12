package tribonacci

import "testing"

func TestTribonacci(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{4, 4},
		{25, 1389537},
		{0, 0},
		{1, 1},
		{2, 1},
	}
	for _, test := range tests {
		if got := tribonacci(test.n); got != test.want {
			t.Errorf("tribonacci(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
