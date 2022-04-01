package two_sum_2

func rank(xs []int, x int) int {
	lo, hi := 0, len(xs)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if x < xs[mid] {
			hi = mid - 1
		} else if x > xs[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

func binarySearch(xs []int, x int) (int, bool) {
	if i := rank(xs, x); i < len(xs) && xs[i] == x {
		return i, true
	}
	return 0, false
}

func twoSum(numbers []int, target int) []int {
	for i, n := range numbers {
		diff := target - n
		diffIndex, ok := binarySearch(numbers, diff)
		if ok && diffIndex != i {
			return []int{i + 1, diffIndex + 1}
		}
	}
	return nil
}
