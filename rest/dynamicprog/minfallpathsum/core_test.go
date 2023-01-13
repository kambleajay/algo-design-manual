package minfallpathsum

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, 13},
		{[][]int{{-19, 57}, {-40, -5}}, -59},
		{[][]int{{-48}}, -48},
	}
	for _, test := range tests {
		if got := minFallingPathSum(test.matrix); got != test.want {
			t.Errorf("minFallingPathSum(%v) = %d, want %d", test.matrix, got, test.want)
		}
	}
}
