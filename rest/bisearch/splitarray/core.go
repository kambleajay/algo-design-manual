package splitarray

import "math"

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func minSubarraysRequired(maxSumAllowed int, nums []int) int {
	var noOfSplits, currSum int
	for _, n := range nums {
		if currSum+n > maxSumAllowed {
			currSum = n
			noOfSplits++
		} else {
			currSum += n
		}
	}
	return noOfSplits + 1
}

func splitArray(nums []int, k int) int {
	var sum int
	maxElem := math.MinInt64
	for _, n := range nums {
		sum += n
		maxElem = max(maxElem, n)
	}
	lo, hi := maxElem, sum
	var minOfLarrgestSplitSum int
	for lo <= hi {
		maxSumAllowed := lo + (hi-lo)/2
		if minSubarraysRequired(maxSumAllowed, nums) <= k {
			hi = maxSumAllowed - 1
			minOfLarrgestSplitSum = maxSumAllowed
		} else {
			lo = maxSumAllowed + 1
		}
	}
	return minOfLarrgestSplitSum
}
