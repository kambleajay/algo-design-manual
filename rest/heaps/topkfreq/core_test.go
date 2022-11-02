package topkfreq

import (
	"algo/testing/utils"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}
	for _, test := range tests {
		if got := topKFrequent(test.nums, test.k); !utils.ContainsAllIntSlice(got, test.want) {
			t.Errorf("topKFrequent(%v, %d) = %v, want %v", test.nums, test.k, got, test.want)
		}
	}
}
