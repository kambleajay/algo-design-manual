package three_sum

import (
	"algo/testing/utils"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var tests = []struct {
		input []int
		want  [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
		{[]int{}, [][]int{}},
		{[]int{0}, [][]int{}},
		{[]int{-2, 0, 0, 2, 2}, [][]int{[]int{-2, 0, 2}}},
	}
	for _, test := range tests {
		if got := threeSum(test.input); !utils.ContainsAll(test.want, got) {
			t.Errorf("threeSum(%v) should be %v, but was %v", test.input, test.want, got)
		}
	}
}
