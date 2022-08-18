package heapsort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		if got := heapSort(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("heapSort(%v) should be %v, but was %v\n", copyOfInput, test.want, got)
		}
	}
}
