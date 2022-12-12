package robber

import "testing"

func TestRob(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{5}, 5},
		{[]int{}, 0},
	}
	for _, test := range tests {
		if got := rob(test.nums); got != test.want {
			t.Errorf("rob(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}
