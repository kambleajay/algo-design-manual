package median_of_two_sorted_arrays

import "math"

func maxFromLeftHalf(a []int, count int) int {
	if count == 0 {
		return math.MinInt64
	}
	return a[count-1]
}

func minFromRightHalf(a []int, count int, limit int) int {
	if count == limit {
		return math.MaxInt64
	}
	return a[count]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays2(nums2, nums1)
	}
	size := m + n
	lo, hi := 0, m
	for lo <= hi {
		elemsFrom1 := lo + (hi-lo)/2
		elemsFrom2 := (m+n+1)/2 - elemsFrom1
		l1 := maxFromLeftHalf(nums1, elemsFrom1)
		l2 := maxFromLeftHalf(nums2, elemsFrom2)
		r1 := minFromRightHalf(nums1, elemsFrom1, m)
		r2 := minFromRightHalf(nums2, elemsFrom2, n)
		if l1 <= r2 && l2 <= r1 {
			if size%2 == 0 {
				return float64(max(l1, l2)+min(r1, r2)) / 2
			}
			return float64(max(l1, l2))
		} else if l1 > r2 {
			hi = elemsFrom1 - 1
		} else {
			lo = elemsFrom1 + 1
		}
	}
	return 0
}
