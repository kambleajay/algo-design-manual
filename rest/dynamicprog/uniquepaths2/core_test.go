package uniquepaths2

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
		{[][]int{{0, 0}, {0, 1}}, 0},
		{[][]int{{1}}, 0},
		{[][]int{{0}}, 1},
	}
	for _, test := range tests {
		if got := uniquePathsWithObstacles(test.grid); got != test.want {
			t.Errorf("uniquePathsWithObstacles(%v) = %d, want %d", test.grid, got, test.want)
		}
	}
}
