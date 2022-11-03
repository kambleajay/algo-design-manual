package kthlarstream

import "testing"

func TestKthLargest(t *testing.T) {
	kthlarg := Constructor(3, []int{4, 5, 8, 2})
	expectations := [][]int{{3, 4}, {5, 5}, {10, 5}, {9, 8}, {4, 8}}
	for _, exp := range expectations {
		if got := kthlarg.Add(exp[0]); got != exp[1] {
			t.Errorf("Add(%d) = %d, want %d", exp[0], got, exp[1])
		}
	}
}
