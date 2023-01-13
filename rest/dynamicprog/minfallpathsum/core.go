package minfallpathsum

import "math"

func min(xs ...int) int {
	answer := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < answer {
			answer = xs[i]
		}
	}
	return answer
}

func minFallingPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for col := 0; col < n; col++ {
		dp[0][col] = matrix[0][col]
	}
	var path1, path2, path3, bestPath int
	for row := 1; row < m; row++ {
		for col := 0; col < n; col++ {
			path1, path2, path3 = math.MaxInt64, math.MaxInt64, math.MaxInt64
			if row-1 >= 0 && col-1 >= 0 {
				path1 = dp[row-1][col-1]
			}
			if row-1 >= 0 {
				path2 = dp[row-1][col]
			}
			if row-1 >= 0 && col+1 < n {
				path3 = dp[row-1][col+1]
			}
			bestPath = min(path1, path2, path3)
			dp[row][col] = bestPath + matrix[row][col]
		}
	}
	answer := math.MaxInt64
	for col := 0; col < n; col++ {
		answer = min(answer, dp[m-1][col])
	}
	return answer
}
