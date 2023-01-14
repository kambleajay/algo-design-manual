package painthouse

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dpfn(i int, lastColor int, costs [][]int) int {
	answer := math.MaxInt64
	if i == len(costs) {
		return 0
	}
	for color, cost := range costs[i] {
		if color != lastColor {
			answer = min(answer, dpfn(i+1, color, costs)+cost)
		}
	}
	return answer
}

func minCost(costs [][]int) int {
	noOfHouses, noOfColors := len(costs), len(costs[0])
	dp := make([][]int, noOfHouses+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, noOfColors)
	}
	for house := noOfHouses - 1; house >= 0; house-- {
		for color, cost := range costs[house] {
			answer := math.MaxInt64
			for diffColor := 0; diffColor < noOfColors; diffColor++ {
				if diffColor != color {
					answer = min(answer, dp[house+1][diffColor]+cost)
				}
			}
			dp[house][color] = answer
		}
	}
	answer := math.MaxInt64
	for _, cost := range dp[0] {
		answer = min(answer, cost)
	}
	return answer
}
