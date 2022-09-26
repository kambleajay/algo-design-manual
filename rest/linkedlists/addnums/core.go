package addnums

import ll "algo/rest/linkedlists"

func addTwoNumbers(l1 *ll.ListNode, l2 *ll.ListNode) *ll.ListNode {
	head := &ll.ListNode{}
	tail := head
	p1, p2 := l1, l2
	var sum, carry, n1, n2 int
	for p1 != nil || p2 != nil || carry != 0 {
		if p1 != nil {
			n1 = p1.Val
		} else {
			n1 = 0
		}
		if p2 != nil {
			n2 = p2.Val
		} else {
			n2 = 0
		}
		sum = n1 + n2 + carry
		carry = sum / 10
		tail.Next = &ll.ListNode{Val: sum % 10}
		tail = tail.Next
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	return head.Next
}
