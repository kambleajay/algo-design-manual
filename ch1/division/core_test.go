package division

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestDivide(t *testing.T) {
	var tests = []struct {
		dividend int
		divisor  int
		want     int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{5, 5, 1},
		{-3, 3, -1},
		{-6, -6, 1},
		{-2147483648, -1, 2147483647},
		{3, 1, 3},
		{10, 1, 10},
		{2147483647, 1, 2147483647},
		{2147483647, -1, -2147483647},
		{13, 2, 6},
		{2147483647, 2, 1073741823},
		{-1, 1, -1},
		{15, 2, 7},
	}
	for _, test := range tests {
		if got := divide(test.dividend, test.divisor); got != test.want {
			t.Errorf("%d/%d != %d (%d)\n", test.dividend, test.divisor, test.want, got)
		}
	}
}

func BenchmarkDivide(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		dividend := rng.Intn(math.MaxInt32-math.MinInt32) + math.MinInt32
		divisor := rng.Intn(math.MaxInt32-math.MinInt32) + math.MinInt32
		divide(dividend, divisor)
	}
}
