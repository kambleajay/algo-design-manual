package reverse

import (
	ll "algo/rest/linkedlists"
	"algo/rest/linkedlists/utils"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	var tests = []struct {
		list *ll.ListNode
		want []int
	}{
		{utils.ToList(1, 2, 3, 4, 5), []int{5, 4, 3, 2, 1}},
		{utils.ToList(1, 2), []int{2, 1}},
		{utils.ToList(), nil},
	}
	for _, test := range tests {
		if got := reverseList(test.list); !reflect.DeepEqual(utils.ToSlice(got), test.want) {
			t.Errorf("reverseList(%v) = %v, want %v", test.list, got, test.want)
		}
	}
}
