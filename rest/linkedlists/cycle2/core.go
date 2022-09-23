package cycle2

type ListNode struct {
	Val  int
	Next *ListNode
}

func intersection(head *ListNode) *ListNode {
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

func detectCycle(head *ListNode) *ListNode {
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
