package coinchange2

func change(amount int, coins []int) int {
	dp := make([][]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(coins)+1)
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = 1 //when amount is 0, there is 1 way to make it
	}
	var withoutCurrCoin, withCurrCoin, amountAfterUsingCurrCoin int
	for j := len(dp[0]) - 2; j >= 0; j-- {
		for i := 1; i < len(dp); i++ {
			withoutCurrCoin = dp[i][j+1]
			amountAfterUsingCurrCoin = i - coins[j]
			if amountAfterUsingCurrCoin < 0 {
				withCurrCoin = 0
			} else {
				withCurrCoin = dp[i-coins[j]][j]
			}
			dp[i][j] = withoutCurrCoin + withCurrCoin
		}
	}
	return dp[amount][0]
}
