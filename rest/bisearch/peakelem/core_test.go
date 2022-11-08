package peakelem

import (
	"algo/testing/utils"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 1}, []int{2}},
		{[]int{1, 2, 1, 3, 5, 6, 4}, []int{1, 5}},
	}
	for _, test := range tests {
		if got := findPeakElement(test.nums); utils.NotIn(got, test.want) {
			t.Errorf("findPeakElement(%v) =  %d, want one of %v", test.nums, got, test.want)
		}
	}
}
