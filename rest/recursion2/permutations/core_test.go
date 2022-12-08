package permutations

import (
	"reflect"
	"testing"
)

func in(x []int, ys [][]int) bool {
	for _, y := range ys {
		if reflect.DeepEqual(x, y) {
			return true
		}
	}
	return false
}

func ContainsAll(xs, ys [][]int) bool {
	for _, x := range xs {
		if !in(x, ys) {
			return false
		}
	}
	return true
}

func TestPermute(t *testing.T) {
	var tests = []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{[]int{1}, [][]int{{1}}},
	}
	for _, test := range tests {
		if got := permute(test.nums); !ContainsAll(got, test.want) {
			t.Errorf("permute(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}
