package maxlenrepsuba

import "testing"

func TestFindLength(t *testing.T) {
	var tests = []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	}
	for _, test := range tests {
		if got := findLength(test.nums1, test.nums2); got != test.want {
			t.Errorf("findLength(%v, %v) = %d, want %d", test.nums1, test.nums2, got, test.want)
		}
	}
}
