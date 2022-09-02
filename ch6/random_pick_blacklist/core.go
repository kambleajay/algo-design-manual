package random_pick_blacklist

import (
	"math/rand"
	"time"
)

type Solution struct {
	n         int
	nums      []int
	size      int
	blacklist []int
	rng       *rand.Rand
}

func Constructor(n int, blacklist []int) Solution {
	allNums := make([]int, n)
	for _, blacklistNum := range blacklist {
		allNums[blacklistNum] = 1
	}
	var nums []int
	for i, num := range allNums {
		if num == 0 {
			nums = append(nums, i)
		}
	}
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	return Solution{n, nums, len(nums), blacklist, rng}
}

func (this *Solution) Pick() int {
	i := this.rng.Intn(this.size)
	return this.nums[i]
}
