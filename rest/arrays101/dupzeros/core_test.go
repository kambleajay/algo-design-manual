package dupzeros

import (
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{8, 4, 5, 0, 0, 0, 0, 7}, []int{8, 4, 5, 0, 0, 0, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{nil, nil},
		{[]int{1, 5, 2, 0, 6, 8, 0, 6, 0}, []int{1, 5, 2, 0, 0, 6, 8, 0, 0}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		if duplicateZeros(test.input); !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("duplicateZeros(%v) = %v, want %v", copyOfInput, test.input, test.want)
		}
	}
}
