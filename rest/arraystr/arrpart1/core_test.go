package arrpart1

import (
	"math/rand"
	"testing"
	"time"
)

func TestArrayPairSum(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1, 4, 3, 2}, 4},
		{[]int{6, 2, 6, 5, 1, 2}, 9},
		{[]int{1, 1}, 1},
	}
	for _, test := range tests {
		if got := arrayPairSum(test.nums); got != test.want {
			t.Errorf("arrayPairSum(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}

func benchmarkArrayPairSum(b *testing.B, size int) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	nums := make([]int, 2*size)
	limit := int(1e4)
	min, max := -limit, limit
	for i := 0; i < len(nums); i++ {
		nums[i] = rng.Intn(max-min) + min
	}
	for i := 0; i < b.N; i++ {
		arrayPairSum(nums)
	}
}

func Benchmark10(b *testing.B) {
	benchmarkArrayPairSum(b, 10)
}

func Benchmark100(b *testing.B) {
	benchmarkArrayPairSum(b, 100)
}

func Benchmark1000(b *testing.B) {
	benchmarkArrayPairSum(b, 1000)
}

func Benchmark10000(b *testing.B) {
	benchmarkArrayPairSum(b, 10000)
}
