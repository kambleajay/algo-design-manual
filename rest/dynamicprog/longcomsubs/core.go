package longcomsubs

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for j := n - 1; j >= 0; j-- {
		for i := m - 1; i >= 0; i-- {
			if text2[j] == text1[i] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[0][0]
}
