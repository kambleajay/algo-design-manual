package kclosest

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	var tests = []struct {
		nums []int
		k, x int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, 4, []int{3, 4, 5}},
		{[]int{2, 4, 6, 8, 10}, 1, 6, []int{6}},
		{[]int{1, 2, 3}, 3, 10, []int{1, 2, 3}},
	}
	for _, test := range tests {
		if got := findClosestElements(test.nums, test.k, test.x); !reflect.DeepEqual(got, test.want) {
			t.Errorf("findClosestElements(%v, %d, %d) = %v, want %v", test.nums, test.k, test.x, got, test.want)
		}
	}
}
