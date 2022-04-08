package sort_list

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMinPQ(t *testing.T) {
	pq := NewMinPQ()
	pq.insert(10)
	pq.insert(15)
	pq.insert(5)
	if got := pq.delMin(); 5 != got {
		t.Errorf("min should be 5, but was %d", got)
	}
	if got := pq.delMin(); 10 != got {
		t.Errorf("min should be 10, but was %d", got)
	}
	if got := pq.delMin(); 15 != got {
		t.Errorf("min should be 15, but was %d", got)
	}
	pq.insert(7)
	pq.insert(3)
	pq.insert(5)
	if got := pq.delMin(); 3 != got {
		t.Errorf("min should be 3, but was %d", got)
	}
	pq.insert(11)
	pq.insert(9)
	if got := pq.delMin(); 5 != got {
		t.Errorf("min should be 10, but was %d", got)
	}
	pq.insert(13)
	pq.insert(1)
	if got := pq.delMin(); 1 != got {
		t.Errorf("min should be 1, but was %d", got)
	}
}

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

func TestSortList(t *testing.T) {
	var tests = []struct {
		input *ListNode
		want  *ListNode
	}{
		{makeList([]int{4, 2, 1, 3}), makeList([]int{1, 2, 3, 4})},
		{makeList([]int{-1, 5, 3, 4, 0}), makeList([]int{-1, 0, 3, 4, 5})},
		{makeList([]int{}), makeList([]int{})},
	}
	for _, test := range tests {
		if got := sortList(test.input); !reflect.DeepEqual(test.want, got) {
			t.Errorf("sortList(%v) should be %v, but was %v", test.input, test.want, got)
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
		sortList(makeList(input))
	}
}

func Benchmark10(b *testing.B)     { benchmark(b, 10) }
func Benchmark100(b *testing.B)    { benchmark(b, 100) }
func Benchmark1000(b *testing.B)   { benchmark(b, 1000) }
func Benchmark10000(b *testing.B)  { benchmark(b, 10000) }
func Benchmark100000(b *testing.B) { benchmark(b, 100000) }
