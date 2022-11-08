package mininrotarr

import "testing"

func TestFindMin(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{2, 3, 4, 5, 1}, 1},
		{[]int{2, 3, 4, 5, 1}, 1},
		{[]int{4, 5, 1, 2, 3}, 1},
		{[]int{5, 1, 2, 3, 4}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{5}, 5},
		{[]int{2, 1}, 1},
	}
	for _, test := range tests {
		if got := findMin(test.nums); got != test.want {
			t.Errorf("findMin(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}
