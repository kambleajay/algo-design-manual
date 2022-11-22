package searchmat

import "testing"

func TestSearchMatrix(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
		{[][]int{}, 1, false},
		{[][]int{{}}, 1, false},
		{nil, 1, false},
	}
	for _, test := range tests {
		if got := searchMatrix(test.matrix, test.target); got != test.want {
			t.Errorf("searchMatrix(%v, %d) = %t, want %t", test.matrix, test.target, got, test.want)
		}
	}
}
