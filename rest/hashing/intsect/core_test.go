package intsect

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	var tests = []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, nil},
	}
	for _, test := range tests {
		if got := intersection(test.nums1, test.nums2); !reflect.DeepEqual(got, test.want) {
			t.Errorf("intersection(%v, %v) = %v, want %v", test.nums1, test.nums2, got, test.want)
		}
	}
}
