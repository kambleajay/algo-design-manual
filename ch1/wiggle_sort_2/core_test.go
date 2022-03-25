package wiggle_sort_2

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	copyOfInput := make([]int, len(input))
	copy(copyOfInput, input)
	shuffle(input)
	if reflect.DeepEqual(copyOfInput, input) {
		t.Errorf("%v should not be equal to %v after shuffle", input, copyOfInput)
	}
}

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{3, 1, 5, 2, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 3, 2, 1}, []int{1, 1, 2, 2, 3, 3}},
		{[]int{1, 5, 1, 1, 6, 4}, []int{1, 1, 1, 4, 5, 6}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		if quickSort(test.input); !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("%v should be equal to %v after sort, but was %v", copyOfInput, test.want, test.input)
		}
	}
}

func TestWiggleSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{3, 2, 1}, []int{2, 3, 1}},
		{[]int{1, 2, 1}, []int{1, 2, 1}},
		{[]int{1, 2, 3}, []int{2, 3, 1}},
		{[]int{1, 2, 3, 4}, []int{2, 4, 1, 3}},
		{[]int{1, 1, 3, 3}, []int{1, 3, 1, 3}},
		{[]int{4, 3, 2, 1}, []int{2, 4, 1, 3}},
		{[]int{4, 5, 5, 6}, []int{5, 6, 4, 5}},
		{[]int{1, 5, 1, 1, 6, 4}, []int{1, 6, 1, 5, 1, 4}},
		{[]int{1, 3, 2, 2, 3, 1}, []int{2, 3, 1, 3, 1, 2}},
	}
	for _, test := range tests {
		copyOfInput := make([]int, len(test.input))
		copy(copyOfInput, test.input)
		wiggleSort(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("%v should be %v, but was %v", copyOfInput, test.want, test.input)
		}
	}
}
