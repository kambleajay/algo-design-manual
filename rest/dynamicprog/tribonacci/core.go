package tribonacci

func tribonacci(n int) int {
	n1, n2, n3 := 0, 1, 1
	switch n {
	case 0:
		return n1
	case 1:
		return n2
	case 2:
		return n3
	}
	for i := 3; i < n; i++ {
		next := n1 + n2 + n3
		n1, n2, n3 = n2, n3, next
	}
	return n1 + n2 + n3
}
