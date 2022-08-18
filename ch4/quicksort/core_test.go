package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
		{[]int{503, 87, 512, 61, 908, 170, 897, 275, 653, 426, 154, 509, 612, 677, 765, 703},
			[]int{61, 87, 154, 170, 275, 426, 503, 509, 512, 612, 653, 677, 703, 765, 897, 908}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		if qsort(test.input); !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("qsort(%v) should be %v, but was %v\n", copyOfInput, test.want, test.input)
		}
	}
}
