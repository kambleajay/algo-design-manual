package kthsmalldist

import "testing"

func TestSmallestDistancePair(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 3, 1}, 1, 0},
		{[]int{1, 1, 1}, 2, 0},
		{[]int{1, 6, 1}, 3, 5},
	}
	for _, test := range tests {
		copyOfNums := make([]int, len(test.nums))
		copy(copyOfNums, test.nums)
		if got := smallestDistancePair(test.nums, test.k); got != test.want {
			t.Errorf("smallestDistancePair(%v, %d) = %d, want %d", copyOfNums, test.k, got, test.want)
		}
	}
}
