package remnthfromend

import (
	ll "algo/rest/linkedlists"
	"algo/rest/linkedlists/utils"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var tests = []struct {
		list *ll.ListNode
		n    int
		want []int
	}{
		{utils.ToList(1, 2, 3, 4, 5), 2, []int{1, 2, 3, 5}},
		{utils.ToList(1), 1, nil},
		{utils.ToList(1, 2), 1, []int{1}},
	}
	for _, test := range tests {
		if got := removeNthFromEnd(test.list, test.n); !reflect.DeepEqual(utils.ToSlice(got), test.want) {
			t.Errorf("removeNthFromEnd(%v, %d) = %v, want %v", test.list, test.n, got, test.want)
		}
	}
}
