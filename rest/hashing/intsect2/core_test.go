package intsect2

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	var tests = []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}
	for _, test := range tests {
		if got := intersect(test.nums1, test.nums2); !reflect.DeepEqual(got, test.want) {
			t.Errorf("intersect(%v, %v) = %v, want %v", test.nums1, test.nums2, got, test.want)
		}
	}
}
