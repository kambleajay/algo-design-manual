package oddeven

import ll "algo/rest/linkedlists"

func oddEvenList(head *ll.ListNode) *ll.ListNode {
	if head == nil {
		return nil
	}
	oddHead, oddTail, evenHead, evenTail := head, head, head.Next, head.Next
	for evenTail != nil && evenTail.Next != nil {
		oddTail.Next = evenTail.Next
		oddTail = oddTail.Next
		evenTail.Next = oddTail.Next
		evenTail = evenTail.Next
	}
	oddTail.Next = evenHead
	return oddHead
}
