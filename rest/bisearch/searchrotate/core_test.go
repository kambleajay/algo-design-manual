package searchrotate

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{5, 1, 3}, 5, 0},
	}
	for _, test := range tests {
		if got := search(test.nums, test.target); got != test.want {
			t.Errorf("search(%v, %d) = %d, want %d", test.nums, test.target, got, test.want)
		}
	}
}

func BenchmarkToMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 1; i < 1e6; i++ {
			fmt.Printf("%d ", i)
		}
	}
}
