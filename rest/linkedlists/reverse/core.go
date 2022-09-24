package reverse

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pred *ListNode
	curr := head

	for curr != nil {
		succ := curr.Next
		curr.Next = pred
		pred = curr
		curr = succ
	}

	return pred
}
