package two_sum_2

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{[]int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.nums))
		copy(copyOfInput, test.nums)
		if got := twoSum(test.nums, test.target); !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v, %d) should be %v, but was %v.", copyOfInput, test.target, test.want, got)
		}
	}
}
