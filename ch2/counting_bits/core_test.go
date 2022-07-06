package counting_bits

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestCountBits(t *testing.T) {
	var tests = []struct {
		n    int
		want []int
	}{
		{0, []int{0}},
		{1, []int{0, 1}},
		{2, []int{0, 1, 1}},
		{7, []int{0, 1, 1, 2, 1, 2, 2, 3}},
		{32, []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1}},
	}
	for _, test := range tests {
		if got := countBits2(test.n); !reflect.DeepEqual(got, test.want) {
			t.Errorf("countBits(%d) should be %v, but was %v", test.n, test.want, got)
		}
	}
}

func BenchmarkCountBits(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	min, max := 0, 100000
	n := rng.Intn(max-min) + min
	for i := 0; i < b.N; i++ {
		countBits2(n)
	}
}
