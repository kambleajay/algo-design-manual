package painthouse

import "testing"

func TestMinCost(t *testing.T) {
	var tests = []struct {
		costs [][]int
		want  int
	}{
		{[][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}, 10},
		{[][]int{{7, 6, 2}}, 2},
		{[][]int{{3, 5, 3}, {6, 17, 6}, {7, 13, 18}, {9, 10, 18}}, 26},
	}
	for _, test := range tests[2:] {
		if got := minCost(test.costs); got != test.want {
			t.Errorf("minCost(%v) = %d, want %d", test.costs, got, test.want)
		}
	}
}
