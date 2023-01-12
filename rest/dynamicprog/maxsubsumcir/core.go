package maxsubsumcir

import "math"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSum(nums []int) int {
	bestSum := math.MinInt64
	currSum := 0
	for _, num := range nums {
		currSum = max(currSum+num, num)
		bestSum = max(bestSum, currSum)
	}
	return bestSum
}

func maxSubarraySumCircular(nums []int) int {
	maxSumNormal := maxSum(nums)
	n := len(nums)
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	suffixSum := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixSum += nums[i]
		rightMax[i] = max(rightMax[i+1], suffixSum)
	}
	maxSumCircular := nums[0]
	prefixSum := 0
	for i := 0; i < n; i++ {
		prefixSum += nums[i]
		if i < n-1 {
			maxSumCircular = max(maxSumCircular, prefixSum+rightMax[i+1])
		}
	}
	return max(maxSumNormal, maxSumCircular)
}
