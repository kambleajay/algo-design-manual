package finddup

func findDuplicate(nums []int) int {
	var duplicate int
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		count := 0
		for _, n := range nums {
			if n <= mid {
				count++
			}
		}
		if count > mid {
			duplicate = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return duplicate
}
