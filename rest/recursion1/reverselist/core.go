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

func reverseListIter(head *ListNode) *ListNode {
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

func reverseListRecur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	headNext := reverseListRecur(head.Next)
	head.Next.Next = head
	head.Next = nil
	return headNext
}

func reverseList(head *ListNode) *ListNode {
	return reverseListRecur(head)
}
