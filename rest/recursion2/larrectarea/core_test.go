package larrectarea

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	var tests = []struct {
		heights []int
		want    int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{[]int{6, 7, 5, 2, 4, 5, 9, 3}, 16},
	}
	for _, test := range tests {
		if got := largestRectangleArea(test.heights); got != test.want {
			t.Errorf("largestRectangleArea(%v) = %d, want %d", test.heights, got, test.want)
		}
	}
}
