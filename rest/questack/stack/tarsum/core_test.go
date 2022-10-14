package tarsum

import (
	"math/rand"
	"testing"
	"time"
)

func TestTargetSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
		{[]int{1, 0}, 1, 2},
	}
	for _, test := range tests {
		if got := findTargetSumWays(test.nums, test.target); got != test.want {
			t.Errorf("findTargetSumWays(%v, %d) = %d, want %d", test.nums, test.target, got, test.want)
		}
	}
}

func BenchmarkTargetSum(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	nums := make([]int, 20)
	for i := 0; i < 20; i++ {
		nums[i] = rng.Intn(1000)
	}
	min, max := -1000, 1000
	target := rng.Intn(max-min) + min
	for i := 0; i < b.N; i++ {
		findTargetSumWays(nums, target)
	}
}
