package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3, 4}, 0, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 1, []int{4, 1, 2, 3}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{-1}, 2, []int{-1}},
	}
	for _, test := range tests {
		copyOfNums := make([]int, len(test.nums))
		copy(copyOfNums, test.nums)
		if rotate(test.nums, test.k); !reflect.DeepEqual(test.nums, test.want) {
			t.Errorf("rotate(%v, %d) = %v, want %v", copyOfNums, test.k, test.nums, test.want)
		}
	}
}
