package topkfreq

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{3, 3, 3, 3}, 2, []int{3}},
		{[]int{1, 1, 1, 2, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{3, 0, 1, 0}, 1, []int{0}},
	}
	for _, test := range tests {
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("failed to test topKFrequent(%v, %d): %v\n", test.nums, test.k, p)
				panic(p)
			}
		}()
		got := topKFrequent(test.nums, test.k)
		sort.Slice(got, func(i, j int) bool {
			return got[i] < got[j]
		})
		sort.Slice(test.want, func(i, j int) bool {
			return test.nums[i] < test.nums[j]
		})
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("topKFrequent(%v, %d) = %v, want %v", test.nums, test.k, got, test.want)
		}
	}
}
