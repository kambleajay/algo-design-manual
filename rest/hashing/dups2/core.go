package dups2

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func containsNearbyDuplicate(nums []int, k int) bool {
	ktoi := make(map[int]int)
	for j, n := range nums {
		i, ok := ktoi[n]
		if ok && i != j && abs(i-j) <= k {
			return true
		}
		ktoi[n] = j
	}
	return false
}
