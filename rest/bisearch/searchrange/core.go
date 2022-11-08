package searchrange

func leftBoundary(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if lo < len(nums) && nums[lo] == target {
		return lo
	}
	return -1
}

func rightBoundary(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if hi >= 0 && nums[hi] == target {
		return hi
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	return []int{leftBoundary(nums, target), rightBoundary(nums, target)}
}
