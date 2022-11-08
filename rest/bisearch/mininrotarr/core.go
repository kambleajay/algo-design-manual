package mininrotarr

import "math"

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// array is not rotated
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		// an inflection point
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		// all elements on the left of the inflection point are greater than the first element
		// all elements on the right of the inflection point are less than the first element
		if nums[mid] > nums[0] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return math.MaxInt64
}
