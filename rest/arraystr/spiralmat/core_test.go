package spiralmat

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	var tests = []struct {
		input [][]int
		want  []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, []int{1, 2, 4, 6, 8, 7, 5, 3}},
	}
	for _, test := range tests {
		if got := spiralOrder(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("spiralOrder(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
