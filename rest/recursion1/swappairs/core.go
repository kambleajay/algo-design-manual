package swappairs

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	if ln == nil {
		return ""
	}
	return fmt.Sprintf("%d->%s", ln.Val, ln.Next)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second, third := head, head.Next, head.Next.Next
	second.Next = first
	first.Next = swapPairs(third)
	return second
}
