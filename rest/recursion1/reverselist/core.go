package reverselist

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

func reverseList(head *ListNode) *ListNode {
	var first, second *ListNode
	second = head
	for second != nil {
		third := second.Next
		second.Next = first
		first = second
		second = third
	}
	return first
}
