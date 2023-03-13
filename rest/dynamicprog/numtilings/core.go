package numtilings

const MOD = int(1e9) + 7

func numTilings(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dpUp := make([]int, n+1)
	dpDown := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	dpUp[0], dpUp[1], dpDown[0], dpDown[1] = 0, 0, 0, 0
	for i := 2; i <= n; i++ {
		dpUp[i] = (dpDown[i-1] + dp[i-2]) % MOD
		dpDown[i] = (dpUp[i-1] + dp[i-2]) % MOD
		dp[i] = (dp[i-1] + dp[i-2] + dpUp[i-1] + dpDown[i-1]) % MOD
	}
	return dp[n]
}
