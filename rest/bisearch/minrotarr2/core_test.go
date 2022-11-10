package minrotarr2

import "testing"

func TestFindMin(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 4}, 0},
		{[]int{0, 1, 4, 4, 5, 6, 7}, 0},
		{[]int{1, 3, 5}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
		{[]int{1, 1, 1, 1, 1}, 1},
		{[]int{5}, 5},
		{[]int{1, 0, 1, 1, 1}, 0},
		{[]int{1, 1, 1, 0, 1}, 0},
	}
	for _, test := range tests {
		if got := findMin(test.nums); got != test.want {
			t.Errorf("findMin(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}
