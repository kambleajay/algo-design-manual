package buysellstockfee

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

const (
	NotHolding = 0
	Holding    = 1
)

func dpfn(i int, holding int, prices []int, fee int) int {
	if i == len(prices) {
		return 0
	}
	var doNothing, doSomething int
	doNothing = dpfn(i+1, holding, prices, fee)
	if holding == NotHolding {
		doSomething = dpfn(i+1, Holding, prices, fee) - prices[i] //buy stock
	} else {
		doSomething = dpfn(i+1, NotHolding, prices, fee) + prices[i] - fee //sell stock
	}
	return max(doNothing, doSomething)
}

func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	var doNothing, doSomething int
	holdingOptions := []int{NotHolding, Holding}
	for i := len(dp) - 2; i >= 0; i-- {
		for _, j := range holdingOptions {
			doNothing = dp[i+1][j]
			if j == NotHolding {
				doSomething = dp[i+1][Holding] - prices[i]
			} else {
				doSomething = dp[i+1][NotHolding] + prices[i] - fee
			}
			dp[i][j] = max(doNothing, doSomething)
		}
	}
	return dp[0][0]
}
