package kclosepoints

import (
	"algo/testing/utils"
	"testing"
)

func TestKClosest(t *testing.T) {
	var tests = []struct {
		points [][]int
		k      int
		want   [][]int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
		{[][]int{{3, 3}}, 1, [][]int{{3, 3}}},
		{[][]int{{3, 3}, {3, 3}}, 1, [][]int{{3, 3}}},
		{[][]int{{3, 3}, {2, 1}}, 1, [][]int{{2, 1}}},
	}
	for _, test := range tests {
		if got := kClosest(test.points, test.k); !utils.ContainsAll(got, test.want) {
			t.Errorf("kClosest(%v, %d) = %v, want %v", test.points, test.k, got, test.want)
		}
	}
}
