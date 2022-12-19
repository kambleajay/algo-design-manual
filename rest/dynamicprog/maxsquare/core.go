package maxsquare

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(xs ...int) int {
	min := xs[0]
	for _, x := range xs[1:] {
		if x < min {
			min = x
		}
	}
	return min
}

func maximalSquare(matrix [][]byte) int {
	var maxSqSide int
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for row := 1; row < m+1; row++ {
		for col := 1; col < n+1; col++ {
			if matrix[row-1][col-1] == '1' {
				dp[row][col] = 1 + min(dp[row][col-1], dp[row-1][col], dp[row-1][col-1])
				maxSqSide = max(maxSqSide, dp[row][col])
			}
		}
	}
	return maxSqSide * maxSqSide
}
