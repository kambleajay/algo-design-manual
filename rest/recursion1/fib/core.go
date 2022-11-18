package fib

func fibHelper(n int, cache map[int]int) int {
	answer, ok := cache[n]
	if ok {
		return answer
	}
	f := fibHelper(n-1, cache) + fibHelper(n-2, cache)
	cache[n] = f
	return f
}

func fib(n int) int {
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	return fibHelper(n, cache)
}
