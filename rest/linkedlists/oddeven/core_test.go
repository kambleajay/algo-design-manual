package oddeven

import (
	"algo/rest/linkedlists/utils"
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
		{[]int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
	}
	for _, test := range tests {
		if got := oddEvenList(utils.ToList(test.input...)); !reflect.DeepEqual(utils.ToSlice(got), test.want) {
			t.Errorf("oddEvenList(%v) = %v, want %v", test.input, utils.ToSlice(got), test.want)
		}
	}
}
