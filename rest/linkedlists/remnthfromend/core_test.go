package remnthfromend

import (
	"reflect"
	"testing"
)

func toListIter(nodes []int, curr *ListNode) *ListNode {
	if len(nodes) == 0 {
		return curr
	}
	node := &ListNode{Val: nodes[0]}
	curr.Next = node
	return toListIter(nodes[1:], node)
}

func toList(nodes ...int) *ListNode {
	head := &ListNode{}
	toListIter(nodes, head)
	return head.Next
}

func toSlice(head *ListNode) []int {
	var nodes []int
	x := head
	for x != nil {
		nodes = append(nodes, x.Val)
		x = x.Next
	}
	return nodes
}

func TestRemoveNthFromEnd(t *testing.T) {
	var tests = []struct {
		list *ListNode
		n    int
		want []int
	}{
		{toList(1, 2, 3, 4, 5), 2, []int{1, 2, 3, 5}},
		{toList(1), 1, nil},
		{toList(1, 2), 1, []int{1}},
	}
	for _, test := range tests {
		if got := removeNthFromEnd(test.list, test.n); !reflect.DeepEqual(toSlice(got), test.want) {
			t.Errorf("removeNthFromEnd(%v, %d) = %v, want %v", test.list, test.n, got, test.want)
		}
	}
}
