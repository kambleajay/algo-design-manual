package buysellstock1

import "math"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	var largestDifference int
	lowestBuyPrice := math.MaxInt64
	for _, price := range prices {
		if price < lowestBuyPrice {
			lowestBuyPrice = price
		} else {
			largestDifference = max(largestDifference, price-lowestBuyPrice)
		}
	}
	return largestDifference
}
