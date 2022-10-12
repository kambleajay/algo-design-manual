package recursion

func perfSquaresUpto(n int) []int {
	var result []int
	for i := 1; i*i <= n; i++ {
		result = append(result, i*i)
	}
	return result
}

func in(n int, set []int) bool {
	for _, elem := range set {
		if n == elem {
			return true
		}
	}
	return false
}

func dividedBy(n int, count int, squares []int) bool {
	if count == 1 {
		return in(n, squares)
	}
	for _, sq := range squares {
		if dividedBy(n-sq, count-1, squares) {
			return true
		}
	}
	return false
}

func numSquares(n int) int {
	squares := perfSquaresUpto(n)
	count := 1
	for {
		if dividedBy(n, count, squares) {
			return count
		}
		count++
	}
}
