package movezeros

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{0, 0, 0, 0, 1}, []int{1, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		moveZeroes(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("moveZeroes(%v) = %v, want %v", copyOfInput, test.input, test.want)
		}
	}
}
