package merge_two_lists

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for head != nil {
		buf.WriteString(fmt.Sprintf("%d ", head.Val))
		head = head.Next
	}
	buf.WriteString("]")
	return buf.String()
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}
	return head.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return mergeTwoLists1(list1, list2)
}
