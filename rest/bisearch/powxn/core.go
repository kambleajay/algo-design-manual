package powxn

func isEven(x int) bool {
	return x%2 == 0
}

func square(x float64) float64 {
	return x * x
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if isEven(n) {
		return square(myPow(x, n/2))
	}
	return x * myPow(x, n-1)
}
