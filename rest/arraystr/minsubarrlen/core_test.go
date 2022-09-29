package minsubarrlen

import "testing"

func TestMinSubarrayLen(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		{[]int{1, 4, 4}, 4, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
	}
	for _, test := range tests {
		if got := minSubArrayLen(test.target, test.nums); got != test.want {
			t.Errorf("minSubArrayLen(%v, %d) = %d, want %d", test.nums, test.target, got, test.want)
		}
	}
}
