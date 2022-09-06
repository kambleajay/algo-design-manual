package removedups

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		input  []int
		want   int
		result []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1, 1, 1, 1, 1}, 1, []int{1}},
		{[]int{}, 0, []int{}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		got := removeDuplicates(test.input)
		if got != test.want {
			t.Errorf("removeDuplicates(%v) = %d, want %d", copyOfInput, got, test.want)
		}
		if !reflect.DeepEqual(test.input[:got], test.result) {
			t.Errorf("the input array = %v, want %v", test.input, test.result)
		}
	}
}
