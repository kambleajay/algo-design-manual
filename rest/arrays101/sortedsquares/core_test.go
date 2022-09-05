package sortedsquares

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{}, []int{}},
		{nil, nil},
	}
	for _, test := range tests {
		if got := sortedSquares(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("sortedSquares(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
