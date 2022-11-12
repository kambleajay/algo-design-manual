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
		mid1 := int((uint32(lo) + uint32(hi)) >> 1)
		mid2 := (m+n+1)/2 - mid1
		left1 := maxFromLeftHalf(nums1, mid1)
		left2 := maxFromLeftHalf(nums2, mid2)
		right1 := minFromRightHalf(nums1, mid1, m)
		right2 := minFromRightHalf(nums2, mid2, n)
		if left1 <= right2 && left2 <= right1 {
			if size%2 == 0 {
				return float64(max(left1, left2)+min(right1, right2)) / 2
			}
			return float64(max(left1, left2))
		} else if left1 > right2 {
			hi = mid1 - 1
		} else {
			lo = mid1 + 1
		}
	}
	return 0
}
