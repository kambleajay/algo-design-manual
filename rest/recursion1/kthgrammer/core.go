package kthgrammer

import "math"

func complement(x int) int {
	if x == 0 {
		return 1
	}
	return 0
}

func byteAt(n, k int) int {
	if n == 1 {
		return 0
	}
	var prev int
	size := int(math.Pow(2, float64(n-1)))
	if k > size/2 {
		prev = complement(byteAt(n-1, k-(size/2)))
	} else {
		prev = byteAt(n-1, k)
	}
	return prev
}

func kthGrammar(n int, k int) int {
	return byteAt(n, k)
}
