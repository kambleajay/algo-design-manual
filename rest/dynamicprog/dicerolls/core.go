package dicerolls

const N = 1000000007

type Tuple struct {
	n      int
	target int
}

func numRollsToTargetRecur(n int, k int, target int, memo map[Tuple]int) int {
	t := Tuple{n, target}
	ans, ok := memo[t]
	if ok {
		return ans
	}
	if n == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	for i := 1; i <= k; i++ {
		ans += numRollsToTargetRecur(n-1, k, target-i, memo)
		ans %= N
	}
	memo[t] = ans
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func numRollsToIter(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	for diceRemaining := 1; diceRemaining < len(dp); diceRemaining++ {
		for sum := 1; sum < len(dp[0]); sum++ {
			for got := 1; got <= min(sum, k); got++ {
				dp[diceRemaining][sum] += dp[diceRemaining-1][sum-got]
				dp[diceRemaining][sum] %= N
			}
		}
	}
	return dp[n][target]
}

func numRollsToTarget(n int, k int, target int) int {
	//memo := make(map[Tuple]int)
	//return numRollsToTargetRecur(n, k, target, memo)
	return numRollsToIter(n, k, target)
}
