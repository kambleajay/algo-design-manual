package division

import (
	"math"
)

func divide(dividend int, divisor int) int {
	var countOfNeg, answer int
	absDividend, absDivisor := dividend, divisor

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if dividend < 0 {
		absDividend = -absDividend
		countOfNeg++
	}
	if divisor < 0 {
		absDivisor = -absDivisor
		countOfNeg++
	}

	tmpDividend, tmpDivisor := absDividend, absDivisor
	for tmpDividend >= tmpDivisor {
		tmpAnswer := 1
		for tmpDividend >= (tmpDivisor << 1) {
			tmpDivisor <<= 1
			tmpAnswer <<= 1
		}
		tmpDividend -= tmpDivisor
		tmpDivisor = absDivisor
		answer += tmpAnswer
	}

	if countOfNeg == 1 {
		return -answer
	}
	return answer
}
