package longincsubs

import "testing"

func TestLengthOfLIS(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}
	for _, test := range tests {
		if got := lengthOfLIS(test.nums); got != test.want {
			t.Errorf("lengthOfLIS(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}
