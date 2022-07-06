package wiggle_sort_2

import (
	"math/rand"
	"time"
)

func shuffle2(xs []int) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < len(xs); i++ {
		r := rng.Intn(i + 1)
		xs[i], xs[r] = xs[r], xs[i]
	}
}

func partition2(a []int, lo, hi int) int {
	i, j := lo, hi+1
	for {
		for i++; a[i] < a[lo]; i++ {
			if i == hi {
				break
			}
		}
		for j--; a[j] > a[lo]; j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

func qselect(a []int, k int) int {
	lo, hi := 0, len(a)-1
	var j int
	for hi > lo {
		j = partition2(a, lo, hi)
		if j < k {
			lo = j + 1
		} else if j > k {
			hi = j - 1
		} else {
			return a[k]
		}
	}
	return a[k]
}

func mapIndex(i, n int) int {
	return (1 + (2 * i)) % (n | 1)
}

func wiggleSort2(nums []int) {
	N := len(nums)
	if nums == nil || N == 0 || N == 1 {
		return
	}
	lo, hi := 0, N-1
	mid := lo + (hi-lo)/2
	shuffle2(nums)
	median := qselect(nums, mid)
	i, lt, gt := 0, 0, N-1
	for i <= gt {
		if nums[mapIndex(i, N)] > median {
			nums[mapIndex(lt, N)], nums[mapIndex(i, N)] = nums[mapIndex(i, N)], nums[mapIndex(lt, N)]
			lt++
			i++
		} else if nums[mapIndex(i, N)] < median {
			nums[mapIndex(i, N)], nums[mapIndex(gt, N)] = nums[mapIndex(gt, N)], nums[mapIndex(i, N)]
			gt--
		} else {
			i++
		}
	}
}
