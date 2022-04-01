package two_sum

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
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{4, -2, 5, 0, 6, 3, 2, 7}, 1, []int{1, 5}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.nums))
		copy(copyOfInput, test.nums)
		if got := twoSum(test.nums, test.target); !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v, %d) should be %v, but was %v.", copyOfInput, test.target, test.want, got)
		}
	}
}
