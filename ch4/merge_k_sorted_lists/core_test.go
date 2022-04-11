package merge_k_sorted_lists

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func makeList(xs []int) *ListNode {
	if len(xs) == 0 {
		return nil
	}
	head := &ListNode{Val: xs[0]}
	node := head
	for _, x := range xs[1:] {
		node.Next = &ListNode{Val: x}
		node = node.Next
	}
	return head
}

func makeSliceOfList(xs [][]int) []*ListNode {
	var res []*ListNode
	for _, x := range xs {
		res = append(res, makeList(x))
	}
	return res
}

func TestMergeKLists(t *testing.T) {
	var tests = []struct {
		input []*ListNode
		want  *ListNode
	}{
		{makeSliceOfList([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}), makeList([]int{1, 1, 2, 3, 4, 4, 5, 6})},
		{makeSliceOfList([][]int{}), makeList([]int{})},
		{makeSliceOfList([][]int{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}}), makeList([]int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4})},
	}
	for _, test := range tests {
		if got := mergeKLists(test.input); !reflect.DeepEqual(test.want, got) {
			t.Errorf("mergeKLists(%v) should be %v, but was %v", test.input, test.want, got)
		}
	}
}

func benchmark(b *testing.B, size int) {
	var input [][]int
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	min, max := -10000, 10000
	for i := 0; i < size; i++ {
		input = make([][]int, size)
		k := rng.Intn(500 + 1)
		input[i] = make([]int, k)
		for j := 0; j < k; j++ {
			input[i][j] = rng.Intn(max-min) + min
		}
	}
	for i := 0; i < b.N; i++ {
		mergeKLists(makeSliceOfList(input))
	}
}

func Benchmark10(b *testing.B)     { benchmark(b, 10) }
func Benchmark100(b *testing.B)    { benchmark(b, 100) }
func Benchmark1000(b *testing.B)   { benchmark(b, 1000) }
func Benchmark10000(b *testing.B)  { benchmark(b, 10000) }
func Benchmark100000(b *testing.B) { benchmark(b, 100000) }
