package random_pick_blacklist

import (
	"math/rand"
	"time"
)

type Solution1 struct {
	size int
	m    map[int]int
	rng  *rand.Rand
}

func Constructor1(n int, blacklist []int) Solution1 {
	right := make(map[int]bool)
	wlen := n - len(blacklist)
	for i := wlen; i < n; i++ {
		right[i] = true
	}
	for _, b := range blacklist {
		delete(right, b)
	}
	var rightKeys []int
	for k, _ := range right {
		rightKeys = append(rightKeys, k)
	}
	m := make(map[int]int)
	i := 0
	for _, b := range blacklist {
		if b < wlen {
			m[b] = rightKeys[i]
			i++
		}
	}
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	return Solution1{wlen, m, rng}
}

func (this *Solution1) Pick1() int {
	i := this.rng.Intn(this.size)
	k, ok := this.m[i]
	if !ok {
		return i
	}
	return k
}
