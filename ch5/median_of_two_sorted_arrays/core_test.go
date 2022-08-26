package median_of_two_sorted_arrays

import (
	"math/rand"
	"testing"
	"time"
)

func TestMedianOfSortedArrays(t *testing.T) {
	var tests = []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{7, 12, 14, 15}, []int{1, 2, 3, 4, 9, 11}, 8},
		{[]int{3}, []int{}, 3.0},
		{[]int{3}, nil, 3.0},
		{[]int{}, []int{5}, 5.0},
		{[]int{}, []int{1}, 1.0},
		{nil, []int{5}, 5.0},
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2, 3}, []int{2, 3}, 2.0},
		{[]int{1, 2, 3}, []int{4, 5}, 3.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 2.0},
		{[]int{1, 2, 3, 4}, []int{1}, 2.0},
		{[]int{1}, []int{1, 2, 3, 4}, 2.0},
		{[]int{1}, []int{2, 3, 4, 5, 6}, 3.5},
		{[]int{1, 2, 3, 5, 6}, []int{4}, 3.5},
		{[]int{4, 6}, []int{1, 2, 3, 5}, 3.5},
		{[]int{1, 3, 4, 6}, []int{2, 5}, 3.5},
		{[]int{1}, []int{2, 3, 4, 5, 6}, 3.5},
		{[]int{1, 2, 3, 4, 5}, []int{6}, 3.5},
		{[]int{3, 4}, []int{}, 3.5},
		{[]int{}, []int{3, 4}, 3.5},
		{[]int{}, []int{2, 3}, 2.5},
		{[]int{2, 3, 4, 5}, []int{}, 3.5},
	}
	for _, test := range tests {
		if got := findMedianSortedArrays(test.nums1, test.nums2); got != test.want {
			t.Errorf("findMedianSortedArrays(%v, %v) was equal to %f, it should be %f\n", test.nums1, test.nums2, got, test.want)
		}
	}
}

func randomBetween(rng *rand.Rand, min, max int) int {
	return rng.Intn(max-min) + min
}

func BenchmarkMedianOfTwoArrays(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	nums1 := make([]int, randomBetween(rng, 1, 1000))
	nums2 := make([]int, randomBetween(rng, 1, 1000))
	tenTo6 := 1000000
	for i := 0; i < len(nums1); i++ {
		nums1[i] = randomBetween(rng, -tenTo6, tenTo6)
	}
	for i := 0; i < len(nums2); i++ {
		nums2[i] = randomBetween(rng, -tenTo6, tenTo6)
	}
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays2(nums1, nums2)
	}
}
