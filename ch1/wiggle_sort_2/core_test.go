package wiggle_sort_2

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	copyOfInput := make([]int, len(input))
	copy(copyOfInput, input)
	shuffle(input)
	if reflect.DeepEqual(copyOfInput, input) {
		t.Errorf("%v should not be equal to %v after shuffle", input, copyOfInput)
	}
}

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{3, 1, 5, 2, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 3, 2, 1}, []int{1, 1, 2, 2, 3, 3}},
		{[]int{1, 5, 1, 1, 6, 4}, []int{1, 1, 1, 4, 5, 6}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		if quickSort(test.input); !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("%v should be equal to %v after sort, but was %v", copyOfInput, test.want, test.input)
		}
	}
}

func atLeastOneEqual(k []int, a [][]int) bool {
	if k == nil && a == nil {
		return true
	}
	if len(k) == 0 && len(a) == 0 {
		return true
	}
	for _, m := range a {
		if reflect.DeepEqual(k, m) {
			return true
		}
	}
	return false
}

func TestWiggleSort(t *testing.T) {
	var tests = []struct {
		input []int
		wants [][]int
	}{
		{nil, nil},
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 1}, [][]int{{1, 1}}},
		{[]int{1, 2}, [][]int{{1, 2}}},
		{[]int{1, 1, 1}, [][]int{{1, 1, 1}}},
		{[]int{3, 3, 3}, [][]int{{3, 3, 3}}},
		{[]int{3, 2, 1}, [][]int{{2, 3, 1}, {1, 3, 2}}},
		{[]int{1, 2, 1}, [][]int{{1, 2, 1}}},
		{[]int{1, 2, 3}, [][]int{{2, 3, 1}, {1, 3, 2}}},
		{[]int{1, 2, 3, 4}, [][]int{{2, 4, 1, 3}, {1, 3, 2, 4}, {2, 3, 1, 4}, {1, 4, 2, 3}}},
		{[]int{1, 1, 3, 3}, [][]int{{1, 3, 1, 3}}},
		{[]int{4, 3, 2, 1}, [][]int{{2, 4, 1, 3}, {1, 3, 2, 4}, {2, 3, 1, 4}, {1, 4, 2, 3}}},
		{[]int{4, 5, 5, 6}, [][]int{{5, 6, 4, 5}}},
		{[]int{1, 5, 1, 1, 6, 4}, [][]int{{1, 6, 1, 5, 1, 4}, {1, 4, 1, 5, 1, 6}, {1, 6, 1, 4, 1, 5},
			{1, 4, 1, 6, 1, 5}, {1, 5, 1, 4, 1, 6}, {1, 5, 1, 6, 1, 4}}},
		{[]int{1, 3, 2, 2, 3, 1}, [][]int{{2, 3, 1, 3, 1, 2}}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		wiggleSort2(test.input)
		if !atLeastOneEqual(test.input, test.wants) {
			t.Errorf("%v should be one of %v, but was %v", copyOfInput, test.wants, test.input)
		}
	}
}

func TestMapIndex(t *testing.T) {
	var tests = []struct {
		i    int
		N    int
		want int
	}{
		{0, 10, 1},
		{1, 10, 3},
		{2, 10, 5},
		{3, 10, 7},
		{4, 10, 9},
		{5, 10, 0},
		{6, 10, 2},
		{7, 10, 4},
		{8, 10, 6},
		{9, 10, 8},
	}
	for _, test := range tests {
		if got := mapIndex(test.i, test.N); got != test.want {
			t.Errorf("mapIndex(%d,%d) should be %d but was %d\n", test.i, test.N, test.want, got)
		}
	}
}

func benchmark(b *testing.B, size int) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	min, max := 1, 5000
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = rng.Intn(max-min) + min
	}
	for i := 0; i < b.N; i++ {
		wiggleSort2(nums)
	}
}

func Benchmark10(b *testing.B)     { benchmark(b, 10) }
func Benchmark100(b *testing.B)    { benchmark(b, 100) }
func Benchmark1000(b *testing.B)   { benchmark(b, 1000) }
func Benchmark10000(b *testing.B)  { benchmark(b, 10000) }
func Benchmark100000(b *testing.B) { benchmark(b, 100000) }
