package cycle2

import ll "algo/rest/linkedlists"

func intersection(head *ll.ListNode) *ll.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

func detectCycle(head *ll.ListNode) *ll.ListNode {
	isec := intersection(head)
	if isec == nil {
		return nil
	}
	i, j := head, isec
	for i != j {
		i = i.Next
		j = j.Next
	}
	return i
}
