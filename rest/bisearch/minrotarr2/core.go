package minrotarr2

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < nums[hi] {
			hi = mid
		} else if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = hi - 1
		}
	}
	return nums[lo]
}
