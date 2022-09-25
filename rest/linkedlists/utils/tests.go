package utils

import "algo/rest/linkedlists"

func toListIter(nodes []int, curr *linkedlists.ListNode) *linkedlists.ListNode {
	if len(nodes) == 0 {
		return curr
	}
	node := &linkedlists.ListNode{Val: nodes[0]}
	curr.Next = node
	return toListIter(nodes[1:], node)
}

func ToList(nodes ...int) *linkedlists.ListNode {
	head := &linkedlists.ListNode{}
	toListIter(nodes, head)
	return head.Next
}

func ToSlice(head *linkedlists.ListNode) []int {
	var nodes []int
	x := head
	for x != nil {
		nodes = append(nodes, x.Val)
		x = x.Next
	}
	return nodes
}
