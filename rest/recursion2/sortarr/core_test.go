package sortarr

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
		{[]int{1, 1, 3, 3, 2, 2}, []int{1, 1, 2, 2, 3, 3}},
		{[]int{}, []int{}},
		{nil, nil},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		if got := sortArray(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("sortArray(%v) = %v, want %v", copyOfInput, got, test.want)
		}
	}
}
