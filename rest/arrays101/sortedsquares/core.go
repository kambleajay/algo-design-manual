package sortedsquares

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(nums []int) []int {
	N := len(nums)
	answer := make([]int, N)
	lo, hi := 0, N-1
	var n int
	for i := N - 1; i >= 0; i-- {
		if abs(nums[lo]) > abs(nums[hi]) {
			n = nums[lo]
			lo++
		} else {
			n = nums[hi]
			hi--
		}
		answer[i] = n * n
	}
	return answer
}
