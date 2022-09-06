package remzeros

import (
	"algo/testing/utils"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		input  []int
		val    int
		want   int
		result []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 4, 0, 3}},
		{[]int{1, 2, 3, 3, 2, 7}, 2, 4, []int{1, 3, 3, 7}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		got := removeElement(test.input, test.val)
		if got != test.want {
			t.Errorf("removeElement(%v, %d) = %d, want %d", copyOfInput, test.val, got, test.want)
		}
		if !utils.ContainsAllIntSlice(test.input[:test.want], test.result) {
			t.Errorf("removeElement(%v, %d) = %v fails contains all test", copyOfInput, test.val, test.input)
		}
	}
}
