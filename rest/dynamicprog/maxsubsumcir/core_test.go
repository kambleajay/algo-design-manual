package maxsubsumcir

import "testing"

func TestMaxSubarraySumCircular(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1, -2, 3, -2}, 3},
		{[]int{5, -3, 5}, 10},
		{[]int{-3, -2, -3}, -2},
		{[]int{5, 5}, 10},
	}
	for _, test := range tests {
		if got := maxSubarraySumCircular(test.nums); got != test.want {
			t.Errorf("maxSubarraySumCircular(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}
