package buysellstock4

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dp(i int, transactionsRemaining int, holding int, prices []int) int {
	if transactionsRemaining == 0 || i == len(prices) {
		return 0
	}
	doNothing := dp(i+1, transactionsRemaining, holding, prices)
	var doSomething int
	if holding == 0 { //we are not holding a stock
		doSomething = -prices[i] + dp(i+1, transactionsRemaining, 1, prices) //buy
	} else { //we are holding a stock
		doSomething = prices[i] + dp(i+1, transactionsRemaining-1, 0, prices) //sell
	}
	return max(doNothing, doSomething)
}

func maxProfit(k int, prices []int) int {
	dp := make([][][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	holding := []int{0, 1}
	var doSomething int
	for i := len(prices) - 1; i >= 0; i-- {
		for tr := 1; tr <= k; tr++ {
			for h := range holding {
				doNothing := dp[i+1][tr][h]
				if h == 0 {
					doSomething = dp[i+1][tr][1] - prices[i]
				} else {
					doSomething = dp[i+1][tr-1][0] + prices[i]
				}
				dp[i][tr][h] = max(doNothing, doSomething)
			}
		}
	}
	return dp[0][k][0]
}
