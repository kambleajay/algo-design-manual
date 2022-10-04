package happynum

func transform(n int) int {
	var sum, rem int
	quot := n
	for quot != 0 {
		quot, rem = quot/10, quot%10
		sum += rem * rem
	}
	return sum
}

func isHappy(n int) bool {
	seen := make(map[int]bool)
	x := n
	for x != 1 && !seen[x] {
		seen[x] = true
		x = transform(x)
	}
	return x == 1
}
