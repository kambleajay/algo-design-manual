package max_subarray

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	currentSum, bestSum := nums[0], nums[0]
	for _, num := range nums[1:] {
		currentSum = max(num, currentSum+num)
		bestSum = max(bestSum, currentSum)
	}
	return bestSum
}
