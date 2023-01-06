package mincostclimbstairs

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCostClimbingStairs1(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[n-1], dp[n-2])
}

func minCostClimbingStairs(cost []int) int {
	back2, back1 := cost[0], cost[1]
	next := math.MaxInt64
	for i := 2; i < len(cost); i++ {
		next = min(back2, back1) + cost[i]
		back2 = back1
		back1 = next
	}
	return min(back1, back2)
}
