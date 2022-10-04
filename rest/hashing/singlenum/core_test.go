package singlenum

import "testing"

func TestSingleNumber(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}
	for _, test := range tests {
		if got := singleNumber(test.nums); got != test.want {
			t.Errorf("singleNumber(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}
