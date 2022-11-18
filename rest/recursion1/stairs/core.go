package stairs

func climbStairsHelper(n int, cache map[int]int) int {
	answer, ok := cache[n]
	if ok {
		return answer
	}
	answer = climbStairsHelper(n-1, cache) + climbStairsHelper(n-2, cache)
	cache[n] = answer
	return answer
}

func climbStairs(n int) int {
	cache := map[int]int{0: 0, 1: 1, 2: 2}
	return climbStairsHelper(n, cache)
}
