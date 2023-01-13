package minpathsum

import "testing"

func TestMinPathSum(t *testing.T) {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
	}
	for _, test := range tests {
		if got := minPathSum(test.grid); got != test.want {
			t.Errorf("minPathSum(%v) = %d, want %d", test.grid, got, test.want)
		}
	}
}
