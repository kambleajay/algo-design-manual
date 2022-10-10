package insdelgetrand

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	keyToIdx map[int]int
	keys     []int
	rng      *rand.Rand
}

func Constructor() RandomizedSet {
	m := make(map[int]int)
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	return RandomizedSet{keyToIdx: m, rng: rng}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.keyToIdx[val]
	if ok {
		return false
	}
	idx := len(this.keys)
	this.keyToIdx[val] = idx
	this.keys = append(this.keys, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.keyToIdx[val]
	if !ok {
		return false
	}
	lastIdx := len(this.keys) - 1
	last := this.keys[lastIdx]
	this.keys[idx], this.keys[lastIdx] = this.keys[lastIdx], this.keys[idx]
	this.keys = this.keys[:lastIdx]
	this.keyToIdx[last] = idx
	delete(this.keyToIdx, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	max := len(this.keys)
	if max == 1 {
		return this.keys[0]
	}
	idx := this.rng.Intn(max)
	return this.keys[idx]
}
