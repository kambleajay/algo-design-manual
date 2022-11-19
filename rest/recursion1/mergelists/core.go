package mergelists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	if ln == nil {
		return "x"
	}
	return fmt.Sprintf("%d->%s", ln.Val, ln.Next)
}

func merge(list1, list2, head *ListNode) {
	if list1 == nil {
		head.Next = list2
		return
	}
	if list2 == nil {
		head.Next = list1
		return
	}
	if list1.Val <= list2.Val {
		head.Next = list1
		merge(list1.Next, list2, head.Next)
	} else {
		head.Next = list2
		merge(list1, list2.Next, head.Next)
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	merge(list1, list2, head)
	return head.Next
}
