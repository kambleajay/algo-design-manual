package foursum2

import "testing"

func TestFourSumCount(t *testing.T) {
	var tests = []struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
		want  int
	}{
		{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{[]int{0}, []int{0}, []int{0}, []int{0}, 1},
		{[]int{1, 2}, []int{1, 2}, []int{-1, -2}, []int{-1, -2}, 6},
	}
	for _, test := range tests {
		if got := fourSumCount(test.nums1, test.nums2, test.nums3, test.nums4); got != test.want {
			t.Errorf("fourSumCount(%v, %v, %v, %v) = %d, want %d", test.nums1, test.nums2, test.nums3, test.nums4, got, test.want)
		}
	}
}
