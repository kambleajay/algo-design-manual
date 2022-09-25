package remelems

import ll "algo/rest/linkedlists"

func removeElements(head *ll.ListNode, val int) *ll.ListNode {
	sentinel := &ll.ListNode{Next: head}
	pred, curr := sentinel, head
	for curr != nil {
		if curr.Val == val {
			pred.Next = curr.Next
		} else {
			pred = curr
		}
		curr = curr.Next
	}

	return sentinel.Next
}
