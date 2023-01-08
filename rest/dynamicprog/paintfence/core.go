package paintfence

func numWays(n int, k int) int {
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}
	back2 := k
	back1 := k * k
	var next int
	for i := 3; i <= n; i++ {
		next = (k - 1) * (back2 + back1)
		back2 = back1
		back1 = next
	}
	return next
}
