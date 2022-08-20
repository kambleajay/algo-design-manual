package reconstruct_queue

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestReconstructQueue(t *testing.T) {
	var tests = []struct {
		input [][]int
		want  [][]int
	}{
		{[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}, [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}},
		{[][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}, [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}}},
		{[][]int{{5, 0}}, [][]int{{5, 0}}},
	}
	for _, test := range tests {
		copyOfInput := make([][]int, len(test.input))
		copy(copyOfInput, test.input)
		if got := reconstructQueue(test.input); !reflect.DeepEqual(test.want, got) {
			t.Errorf("reconstructQueue(%v) should be %v, but was %v", copyOfInput, test.want, got)
		}
	}
}

func benchmark(b *testing.B, size int) {
	var input [][]int
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	hmin, hmax := 0, 1000000
	kmin, kmax := 0, size-1
	input = make([][]int, size)
	for i := 0; i < size; i++ {
		h := rng.Intn(hmax-hmin) + hmin
		k := rng.Intn(kmax-kmin) + kmin
		input[i] = []int{h, k}
	}
	for i := 0; i < b.N; i++ {
		fmt.Printf("input = %v\n", input)
		reconstructQueue(input)
	}
}

func Benchmark10(b *testing.B)    { benchmark(b, 10) }
func Benchmark100(b *testing.B)   { benchmark(b, 100) }
func Benchmark1000(b *testing.B)  { benchmark(b, 1000) }
func Benchmark10000(b *testing.B) { benchmark(b, 2000) }
