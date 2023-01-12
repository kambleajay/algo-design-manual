package uniquepaths2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 && obstacleGrid[i-1][j] != 1 && obstacleGrid[i][j] != 1 {
				dp[i][j] += dp[i-1][j]
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] != 1 && obstacleGrid[i][j] != 1 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
