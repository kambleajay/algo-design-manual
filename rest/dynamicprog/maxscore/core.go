package maxscore

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	for i := m - 1; i >= 0; i-- {
		for left := i; left >= 0; left-- {
			mult := multipliers[i]
			right := n - 1 - (i - left)
			leftPoints := mult*nums[left] + dp[i+1][left+1]
			rightPoints := mult*nums[right] + dp[i+1][left]
			dp[i][left] = max(leftPoints, rightPoints)
		}
	}
	return dp[0][0]
}
