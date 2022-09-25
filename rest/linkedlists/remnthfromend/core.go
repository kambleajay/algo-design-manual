package remnthfromend

import ll "algo/rest/linkedlists"

func removeNthFromEnd(head *ll.ListNode, n int) *ll.ListNode {
	curr := head
	for i := 0; i < n; i++ {
		curr = curr.Next
	}

	if curr == nil {
		return head.Next
	}

	pred := head
	for curr.Next != nil {
		curr = curr.Next
		pred = pred.Next
	}

	pred.Next = pred.Next.Next

	return head
}
