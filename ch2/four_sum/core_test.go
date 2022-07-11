package four_sum

import (
	"algo/testing/utils"
	"testing"
)

func TestFourSum(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
		want   [][]int
	}{
		{[]int{}, 0, [][]int{}},
		{[]int{1, 2, 3}, 0, [][]int{}},
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
	}
	for _, test := range tests[2:3] {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		if got := fourSum(test.input, test.target); !utils.ContainsAll(test.want, got) {
			t.Errorf("fourSum(%v, 4) should be %v, but was %v", copyOfInput, test.want, got)
		}
	}
}
