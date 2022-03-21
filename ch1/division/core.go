package division

import (
	"log"
	"math"
)

func Divide1(dividend, divisor int) int {
	dividendAbs := math.Abs(float64(dividend))
	divisorAbs := math.Abs(float64(divisor))
	result := 1
	for divisorAbs < dividendAbs {
		result += 1
		divisorAbs += divisorAbs
	}
	return result
}

func Divide2(x, y float64) float64 {
	diffOfLogs := math.Log(x) - math.Log(y)
	result := math.Exp(diffOfLogs)
	log.Printf("diff of logs: %f, e^dol = %f\n", diffOfLogs, result)
	return result
}
