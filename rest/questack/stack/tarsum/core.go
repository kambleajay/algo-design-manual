package tarsum

import "math"

func findTargetSumWaysRecur(nums []int, i, sum, target int, seen [][]int, total int) int {
	if i == len(nums) {
		if sum == target {
			return 1
		}
		return 0
	}
	cached := seen[i][sum+total]
	if cached != math.MinInt64 {
		return cached
	}
	num := nums[i]
	answer := findTargetSumWaysRecur(nums, i+1, sum-num, target, seen, total) + findTargetSumWaysRecur(nums, i+1, sum+num, target, seen, total)
	seen[i][sum+total] = answer
	return answer
}

func findTargetSumWays1(nums []int, target int) int {
	var total int
	for _, n := range nums {
		total += n
	}
	seen := make([][]int, len(nums))
	for i := 0; i < len(seen); i++ {
		seen[i] = make([]int, (2*total)+1)
	}
	for i := 0; i < len(seen); i++ {
		for j := 0; j < len(seen[i]); j++ {
			seen[i][j] = math.MinInt64
		}
	}
	return findTargetSumWaysRecur(nums, 0, 0, target, seen, total)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func iterate(from, to int) []int {
	var result []int
	for i := from; i <= to; i++ {
		result = append(result, i)
	}
	return result
}

func findTargetSumWays(nums []int, target int) int {
	var total int
	for _, n := range nums {
		total += n
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, (2*total)+1)
	}
	dp[0][nums[0]+total] = 1
	dp[0][-nums[0]+total] += 1
	for i := 1; i < len(nums); i++ {
		for _, sum := range iterate(-total, total) {
			sumCount := dp[i-1][sum+total]
			if sumCount > 0 {
				dp[i][sum+nums[i]+total] += sumCount
				dp[i][sum-nums[i]+total] += sumCount
			}
		}
	}
	if abs(target) > total {
		return 0
	}
	return dp[len(nums)-1][target+total]
}
