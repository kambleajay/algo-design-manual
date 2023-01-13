package minpathsum

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	var viaLeft, viaTop int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 || j != 0 {
				viaLeft = math.MaxInt64
				viaTop = math.MaxInt64
				if i-1 >= 0 {
					viaLeft = dp[i-1][j]
				}
				if j-1 >= 0 {
					viaTop = dp[i][j-1]
				}
				dp[i][j] = grid[i][j] + min(viaLeft, viaTop)
			}
		}
	}
	return dp[m-1][n-1]
}
