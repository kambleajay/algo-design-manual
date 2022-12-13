package deleteandearn

import "testing"

func TestDeleteAndEarn(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 2}, 6},
		{[]int{2, 2, 3, 3, 3, 4}, 9},
	}
	for _, test := range tests {
		if got := deleteAndEarn(test.nums); got != test.want {
			t.Errorf("deleteAndEarn(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}
