package kthlargest

import "testing"

func TestFindKthLargest(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, test := range tests {
		if got := findKthLargest(test.nums, test.k); got != test.want {
			t.Errorf("findKthLargest(%v, %d) = %d, want %d", test.nums, test.k, got, test.want)
		}
	}
}
