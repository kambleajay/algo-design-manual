package minsubarrlen

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt64
	var sum, left int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			ans = min(ans, i+1-left)
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
