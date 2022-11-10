package finddup

import "testing"

func TestFindDuplicate(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{2, 2, 2}, 2},
		{[]int{2, 2, 2, 2}, 2},
	}
	for _, test := range tests {
		if got := findDuplicate(test.nums); got != test.want {
			t.Errorf("findDuplicate(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}
