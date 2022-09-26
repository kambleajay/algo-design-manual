package addnums

import (
	"algo/rest/linkedlists/utils"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, test := range tests {
		if got := addTwoNumbers(utils.ToList(test.l1...), utils.ToList(test.l2...)); !reflect.DeepEqual(utils.ToSlice(got), test.want) {
			t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", test.l1, test.l2, utils.ToSlice(got), test.want)
		}
	}
}
