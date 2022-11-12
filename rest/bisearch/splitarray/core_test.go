package splitarray

import "testing"

func TestSplitArray(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{7, 2, 5, 10, 8}, 2, 18},
		{[]int{1, 2, 3, 4, 5}, 2, 9},
	}
	for _, test := range tests {
		if got := splitArray(test.nums, test.k); got != test.want {
			t.Errorf("splitArray(%v, %d) = %d, want %d", test.nums, test.k, got, test.want)
		}
	}
}
