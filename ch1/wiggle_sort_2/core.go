package wiggle_sort_2

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

func merge(zs *[]int, xs []int, ys []int) {
	if len(xs) == 0 && len(ys) == 0 {
		return
	}
	var nextXs, nextYs []int
	if len(xs) > 0 {
		*zs = append(*zs, xs[len(xs)-1])
		nextXs = xs[:len(xs)-1]
	}
	if len(ys) > 0 {
		*zs = append(*zs, ys[len(ys)-1])
		nextYs = ys[:len(ys)-1]
	}
	merge(zs, nextXs, nextYs)
}

func wiggleSort(nums []int) {
	if nums == nil || len(nums) == 0 || len(nums) == 1 {
		return
	}
	shuffle(nums)
	quickSort(nums)
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		mid--
	}
	var answer []int
	merge(&answer, nums[:mid+1], nums[mid+1:])
	copy(nums, answer)
}
