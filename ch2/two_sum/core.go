package two_sum

import (
	"math/rand"
	"time"
)

func shuffle(xs []int) {
	if xs == nil {
		return
	}
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < len(xs); i++ {
		r := rng.Intn(i + 1)
		xs[i], xs[r] = xs[r], xs[i]
	}
}

func partition(xs []int, lo int, hi int) int {
	i, j := lo, hi+1
	for {
		for i++; xs[i] < xs[lo]; i++ {
			if i == hi {
				break
			}
		}
		for j--; xs[j] > xs[lo]; j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		xs[i], xs[j] = xs[j], xs[i]
	}
	xs[lo], xs[j] = xs[j], xs[lo]
	return j
}

func quickSortRecur(xs []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := partition(xs, lo, hi)
	quickSortRecur(xs, lo, j-1)
	quickSortRecur(xs, j+1, hi)
}

func quickSort(xs []int) {
	if xs == nil || len(xs) == 1 {
		return
	}
	quickSortRecur(xs, 0, len(xs)-1)
}

func binarySearchRecur(xs []int, x, lo, hi int) (int, bool) {
	if lo > hi {
		return 0, false
	}
	mid := (lo + hi) / 2
	if xs[mid] == x {
		return mid, true
	}
	if xs[mid] > x {
		return binarySearchRecur(xs, x, lo, mid-1)
	} else {
		return binarySearchRecur(xs, x, mid+1, hi)
	}
}

func binarySearch(xs []int, x int) (int, bool) {
	return binarySearchRecur(xs, x, 0, len(xs))
}

func twoSum(nums []int, target int) []int {
	shuffle(nums)
	quickSort(nums)
	for i, n := range nums {
		diff := target - n
		diffIndex, ok := binarySearch(nums, diff)
		if ok && diffIndex != i {
			return []int{i, diffIndex}
		}
	}
	return nil
}
