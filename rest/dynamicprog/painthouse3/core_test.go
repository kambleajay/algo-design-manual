package painthouse3

import (
	"math/rand"
	"testing"
	"time"
)

func TestMinCost(t *testing.T) {
	var tests = []struct {
		houses       []int
		cost         [][]int
		m, n, target int
		want         int
	}{
		{[]int{0, 0, 0, 0, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3, 9},
		{[]int{0, 2, 1, 2, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3, 11},
		{[]int{3, 1, 2, 3}, [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 4, 3, 3, -1},
	}
	for _, test := range tests {
		if got := minCost(test.houses, test.cost, test.m, test.n, test.target); got != test.want {
			t.Errorf("minCost(%v, %v, %d, %d, %d) = %d, want %d", test.houses, test.cost, test.m, test.n, test.target, got, test.want)
		}
	}
}

func BenchmarkMinCost(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	houses := make([]int, 100)
	//for i := 0; i < len(houses); i++ {
	//	houses[i] = rng.Intn(20)
	//}
	noOfColors := 20
	cost := make([][]int, len(houses))
	for i := 0; i < len(cost); i++ {
		cost[i] = make([]int, noOfColors)
		for j := 0; j < len(cost[i]); j++ {
			cost[i][j] = rng.Intn(10000-1) + 1
		}
	}
	target := rng.Intn(100-1) + 1
	for i := 0; i < b.N; i++ {
		minCost(houses, cost, 100, 20, target)
	}
}
