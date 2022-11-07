package search

import "testing"

func TestSearch(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}
	for _, test := range tests {
		if got := search(test.nums, test.target); got != test.want {
			t.Errorf("search(%v, %d) = %d, want %d", test.nums, test.target, got, test.want)
		}
	}
}
