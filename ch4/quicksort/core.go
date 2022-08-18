package quicksort

import (
	"math/rand"
	"time"
)

func less(a []int, i, j int) bool {
	return a[i] < a[j]
}

func exch(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func partition(a []int, lo, hi int) int {
	i, j := lo, hi+1
	for {
		for i++; less(a, i, lo); i++ {
			if i == hi {
				break
			}
		}
		for j--; less(a, lo, j); j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		exch(a, i, j)
	}
	exch(a, lo, j)
	return j
}

func qsortIter(a []int, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	qsortIter(a, lo, j-1)
	qsortIter(a, j+1, hi)

}

func shuffle(a []int) {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < len(a); i++ {
		r := rng.Intn(i + 1)
		a[i], a[r] = a[r], a[i]
	}
}

func qsort(a []int) {
	if a == nil || len(a) == 0 {
		return
	}
	shuffle(a)
	qsortIter(a, 0, len(a)-1)
}
