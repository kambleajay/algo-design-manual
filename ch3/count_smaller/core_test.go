package count_smaller

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestCountSmaller(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{-1}, []int{0}},
		{[]int{-1, -1}, []int{0, 0}},
		{[]int{15, 4, 10, 6, 9, 20, 3}, []int{5, 1, 3, 1, 1, 1, 0}},
		{[]int{5, 17, 11, 10, 16, 15, 13, 1}, []int{1, 6, 2, 1, 3, 2, 1, 0}},
		{[]int{19, 5, 33, 39, 41}, []int{1, 0, 0, 0, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{0, 0, 0, 0, 0}},
		{[]int{83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41}, []int{17, 9, 19, 12, 14, 6, 12, 5, 12, 3, 0, 10, 0, 7, 8, 4, 0, 0, 4, 3, 2, 0, 1, 0}},
		{[]int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41}, []int{10, 27, 10, 35, 12, 22, 28, 8, 19, 2, 12, 2, 9, 6, 12, 5, 17, 9, 19, 12, 14, 6, 12, 5, 12, 3, 0, 10, 0, 7, 8, 4, 0, 0, 4, 3, 2, 0, 1, 0}},
	}
	for _, test := range tests {
		if got := countSmaller(test.input); !reflect.DeepEqual(test.want, got) {
			t.Errorf("countSmaller(%v) should be %v, but was %v", test.input, test.want, got)
		}
	}
}

func benchmark(b *testing.B, size int) {
	var input []int
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	min, max := -10000, 10000
	for i := 0; i < size; i++ {
		input = append(input, rng.Intn(max-min)+min)
	}
	for i := 0; i < b.N; i++ {
		countSmaller(input)
	}
}

func Benchmark10(b *testing.B)     { benchmark(b, 10) }
func Benchmark100(b *testing.B)    { benchmark(b, 100) }
func Benchmark1000(b *testing.B)   { benchmark(b, 1000) }
func Benchmark10000(b *testing.B)  { benchmark(b, 10000) }
func Benchmark100000(b *testing.B) { benchmark(b, 100000) }

func TestRank(t *testing.T) {
	tree := RedBlackBST{}
	items := []int{5, 2, 6, 1}
	for _, item := range items {
		tree.Put(item, item)
	}
	answer := []int{2, 1, 3, 0}
	for i, item := range items {
		if got := tree.Rank(item); answer[i] != got {
			t.Errorf("Rank(%d) should be %d, but was %d", item, answer[i], got)
		}
	}
}
