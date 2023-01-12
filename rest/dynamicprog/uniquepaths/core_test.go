package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	var tests = []struct {
		m, n int
		want int
	}{
		{3, 7, 28},
		{3, 2, 3},
	}
	for _, test := range tests {
		if got := uniquePaths(test.m, test.n); got != test.want {
			t.Errorf("uniquePaths(%d, %d) = %d, want %d", test.m, test.n, got, test.want)
		}
	}
}
