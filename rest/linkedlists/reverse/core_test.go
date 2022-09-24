package reverse

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

func TestReverseList(t *testing.T) {
	var tests = []struct {
		list *ListNode
		want []int
	}{
		{toList(1, 2, 3, 4, 5), []int{5, 4, 3, 2, 1}},
		{toList(1, 2), []int{2, 1}},
		{toList(), nil},
	}
	for _, test := range tests {
		if got := reverseList(test.list); !reflect.DeepEqual(toSlice(got), test.want) {
			t.Errorf("reverseList(%v) = %v, want %v", test.list, got, test.want)
		}
	}
}
