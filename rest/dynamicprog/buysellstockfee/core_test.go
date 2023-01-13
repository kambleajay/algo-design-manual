package buysellstockfee

import (
	"math/rand"
	"testing"
	"time"
)

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		prices []int
		fee    int
		want   int
	}{
		{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
		{[]int{1, 3, 7, 5, 10, 3}, 3, 6},
	}
	for _, test := range tests {
		if got := maxProfit(test.prices, test.fee); got != test.want {
			t.Errorf("maxProfit(%v, %d) = %d, want %d", test.prices, test.fee, got, test.want)
		}
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	max := 5 * 10000
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	prices := make([]int, max)
	for i := 0; i < max; i++ {
		prices[i] = rng.Intn(max-1-1) + 1
	}
	fee := rng.Intn(max)
	for i := 0; i < b.N; i++ {
		maxProfit(prices, fee)
	}
}
