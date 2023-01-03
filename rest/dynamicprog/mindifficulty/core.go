package mindifficulty

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if d > n {
		return -1
	}
	dp := make([][]int, d+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= d; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	var dailyMaxJobDiff int
	for daysRemaining := 1; daysRemaining <= d; daysRemaining++ {
		for i := 0; i < n-daysRemaining+1; i++ {
			dailyMaxJobDiff = 0
			for j := i + 1; j < n-daysRemaining+2; j++ {
				dailyMaxJobDiff = max(dailyMaxJobDiff, jobDifficulty[j-1])
				if dp[daysRemaining-1][j] != math.MaxInt64 {
					dp[daysRemaining][i] = min(dp[daysRemaining][i], dailyMaxJobDiff+dp[daysRemaining-1][j])
				}
			}
		}
	}
	return dp[d][0]
}
