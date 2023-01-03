package coins

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}
	for _, coin := range coins {
		if coin <= amount {
			dp[coin] = 1
		}
	}
	for denom, count := range dp {
		if count != math.MaxInt64 {
			for _, coin := range coins {
				nextDenom := denom + coin
				nextCount := count + 1
				if nextDenom <= amount {
					dp[nextDenom] = min(dp[nextDenom], nextCount)
				}
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}
