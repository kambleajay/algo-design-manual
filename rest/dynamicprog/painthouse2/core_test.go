package painthouse2

import "testing"

func TestMinCost2(t *testing.T) {
	var tests = []struct {
		costs [][]int
		want  int
	}{
		{[][]int{{10, 6, 16, 25, 7, 28}, {7, 16, 18, 30, 16, 25}, {8, 26, 6, 22, 26, 19}, {10, 23, 14, 17, 23, 9}, {12, 14, 27, 7, 8, 9}}, 35},
		{[][]int{{1, 5, 3}, {2, 9, 4}}, 5},
		{[][]int{{1, 3}, {2, 4}}, 5},
	}
	for _, test := range tests {
		if got := minCostII(test.costs); got != test.want {
			t.Errorf("minCostII(%v) = %d, want %d", test.costs, got, test.want)
		}
	}
}
