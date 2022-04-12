package k_pairs_smallest_sums

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestKSmallestPairs(t *testing.T) {
	var tests = []struct {
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{{1, 1}, {1, 1}}},
		{[]int{1, 2}, []int{3}, 3, [][]int{{1, 3}, {2, 3}}},
		{[]int{1}, []int{1, 2, 3}, 2, [][]int{{1, 1}, {1, 2}}},
		{[]int{}, []int{}, 2, [][]int{}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 10, [][]int{{1, 1}, {1, 1}, {2, 1}, {1, 2}, {1, 2}, {2, 2}, {1, 3}, {1, 3}, {2, 3}}},
	}
	for _, test := range tests {
		if got := kSmallestPairs(test.nums1, test.nums2, test.k); !reflect.DeepEqual(test.want, got) {
			t.Errorf("kSmallestPairs(%v, %v, %d) must equal %v, but was %v", test.nums1, test.nums2, test.k, test.want, got)
		}
	}
}

func benchmark(b *testing.B, size int) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	nums1 := make([]int, size)
	nums2 := make([]int, size)
	tenRaiseedToNine := 1000000000
	for i := 0; i < size; i++ {
		nums1[i] = rng.Intn(tenRaiseedToNine-(-tenRaiseedToNine)) + -tenRaiseedToNine
		nums2[i] = rng.Intn(tenRaiseedToNine-(-tenRaiseedToNine)) + -tenRaiseedToNine
	}
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
	sort.Slice(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })
	k := rng.Intn(10000-1) + 1
	for i := 0; i < b.N; i++ {
		kSmallestPairs(nums1, nums2, k)
	}
}

func Benchmark10(b *testing.B)     { benchmark(b, 10) }
func Benchmark100(b *testing.B)    { benchmark(b, 100) }
func Benchmark1000(b *testing.B)   { benchmark(b, 1000) }
func Benchmark10000(b *testing.B)  { benchmark(b, 10000) }
func Benchmark100000(b *testing.B) { benchmark(b, 100000) }
