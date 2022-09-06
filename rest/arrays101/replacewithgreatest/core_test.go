package replacewithgreatest

import (
	"reflect"
	"testing"
)

func TestReplaceElements(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
		{[]int{400}, []int{-1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 5, 5, 5, -1}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		if replaceElements(test.input); !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("replaceElements(%v) = %v, want %v", copyOfInput, test.input, test.want)
		}
	}
}
