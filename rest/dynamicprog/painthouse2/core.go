package painthouse2

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCostIIFirst(costs [][]int) int {
	noOfHouses, noOfColors := len(costs), len(costs[0])
	dp := make([][]int, noOfHouses+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, noOfColors)
	}
	var answer int
	for house := noOfHouses - 1; house >= 0; house-- {
		for color1 := 0; color1 < noOfColors; color1++ {
			answer = math.MaxInt64
			for color2 := 0; color2 < noOfColors; color2++ {
				if color2 != color1 {
					answer = min(answer, dp[house+1][color2]+costs[house][color1])
				}
			}
			dp[house][color1] = answer
		}
	}
	minCost := math.MaxInt64
	for i := 0; i < noOfColors; i++ {
		minCost = min(minCost, dp[0][i])
	}
	return minCost
}

func minCostII(costs [][]int) int {
	m, n := len(costs), len(costs[0])
	for i := 1; i < m; i++ {
		prevRow := costs[i-1]
		min1, min2 := -1, -1
		for k := 0; k < len(prevRow); k++ {
			if min1 == -1 || prevRow[k] < prevRow[min1] {
				min1, min2 = k, min1
			} else if min2 == -1 || prevRow[k] < prevRow[min2] {
				min2 = k
			}
		}
		for j := 0; j < n; j++ {
			if j == min1 {
				costs[i][j] += prevRow[min2]
			} else {
				costs[i][j] += prevRow[min1]
			}
		}
	}
	answer := math.MaxInt64
	lastRow := costs[m-1]
	for _, cost := range lastRow {
		answer = min(answer, cost)
	}
	return answer
}
