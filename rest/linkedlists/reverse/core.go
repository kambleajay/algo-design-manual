package reverse

import ll "algo/rest/linkedlists"

func reverseList(head *ll.ListNode) *ll.ListNode {
	var pred *ll.ListNode
	curr := head

	for curr != nil {
		succ := curr.Next
		curr.Next = pred
		pred = curr
		curr = succ
	}

	return pred
}
