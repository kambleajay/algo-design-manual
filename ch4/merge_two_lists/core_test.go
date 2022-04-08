package merge_two_lists

import (
	"reflect"
	"testing"
)

func makeList(xs []int) *ListNode {
	if len(xs) == 0 {
		return nil
	}
	head := &ListNode{Val: xs[0]}
	node := head
	for _, x := range xs[1:] {
		node.Next = &ListNode{Val: x}
		node = node.Next
	}
	return head
}

func TestMergeTwoLists(t *testing.T) {
	var tests = []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{makeList([]int{1, 2, 4}), makeList([]int{1, 3, 4}), makeList([]int{1, 1, 2, 3, 4, 4})},
		{nil, nil, nil},
		{makeList([]int{}), makeList([]int{0}), makeList([]int{0})},
	}
	for _, test := range tests {
		if got := mergeTwoLists(test.list1, test.list2); !reflect.DeepEqual(test.want, got) {
			t.Errorf("mergeTwoLists(%v, %v) should equal %v, but was %v", test.list1, test.list2, test.want, got)
		}
	}
}
