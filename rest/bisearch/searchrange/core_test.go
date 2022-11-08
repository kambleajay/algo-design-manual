package searchrange

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1, 1, 1, 1, 1}, 1, []int{0, 4}},
	}
	for _, test := range tests {
		if got := searchRange(test.nums, test.target); !reflect.DeepEqual(got, test.want) {
			t.Errorf("searchRange(%v, %d) = %v, want %v", test.nums, test.target, got, test.want)
		}
	}
}
