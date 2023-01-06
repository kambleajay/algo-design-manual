package buysellcooldown

import "math"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

const (
	NotHolding = 0
	Holding    = 1
	NoCooldown = 0
	Cooldown   = 1
)

func dp(i int, holding int, cooldown int, prices []int) int {
	if i == len(prices) {
		return 0
	}
	var doNothing, doSomething int
	if cooldown == Cooldown {
		return dp(i+1, holding, NoCooldown, prices)
	}
	doNothing = dp(i+1, holding, cooldown, prices)
	if holding == NotHolding {
		doSomething = dp(i+1, Holding, cooldown, prices) - prices[i] //buy stock
	} else {
		doSomething = dp(i+1, NotHolding, Cooldown, prices) + prices[i] //sell stock and 1-day cooldown begins
	}
	return max(doNothing, doSomething)
}

func maxProfit1(prices []int) int {
	dp := make([][][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	holding := []int{0, 1}
	cooldown := []int{0, 1}
	var doNothing, doSomething int
	for i := len(prices) - 1; i >= 0; i-- {
		for _, j := range holding {
			for _, k := range cooldown {
				if k == Cooldown {
					dp[i][j][k] = dp[i+1][j][NoCooldown]
				} else {
					doNothing = dp[i+1][j][k]
					if j == NotHolding {
						doSomething = dp[i+1][Holding][k] - prices[i] //buy stock + holding
					} else {
						doSomething = dp[i+1][NotHolding][Cooldown] + prices[i] //sell stock + cooldown
					}
					dp[i][j][k] = max(doNothing, doSomething)
				}
			}
		}
	}
	return dp[0][0][0]
}

func maxProfit2(prices []int) int {
	n := len(prices)
	held := make([]int, n+1)
	sold := make([]int, n+1)
	reset := make([]int, n+1)
	held[0], sold[0] = math.MinInt64, math.MinInt64
	for i := 1; i <= n; i++ {
		held[i] = max(held[i-1], reset[i-1]-prices[i-1])
		sold[i] = held[i-1] + prices[i-1]
		reset[i] = max(reset[i-1], sold[i-1])
	}
	return max(sold[n], reset[n])
}

func maxProfit(prices []int) int {
	return maxProfit2(prices)
}
