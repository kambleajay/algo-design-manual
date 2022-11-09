package perfsqr

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	guess := num / 2
	for (guess * guess) > num {
		guess = (guess + (num / guess)) / 2
	}
	return (guess * guess) == num
}
