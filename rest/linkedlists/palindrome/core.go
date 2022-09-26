package palindrome

import ll "algo/rest/linkedlists"

func endOfFirstHalf(head *ll.ListNode) *ll.ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

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

func isPalindrome(head *ll.ListNode) bool {
	if head == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)
	p1, p2 := head, secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
