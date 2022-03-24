package rotate_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	return fmt.Sprintf("%d->%v", n.Val, n.Next)
}

func nodesSeq(head *ListNode) []*ListNode {
	var nodes []*ListNode
	x := head
	for x != nil {
		nodes = append(nodes, x)
		x = x.Next
	}
	return nodes
}

func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	nodes := nodesSeq(head)

	if len(nodes) == 1 || len(nodes) == k {
		return head
	}

	if k > len(nodes) {
		k %= len(nodes)
	}

	first := nodes[0]
	var secondToLast *ListNode
	for i := k; i > 0; i-- {
		last := nodes[len(nodes)-1]
		if len(nodes) > 2 {
			secondToLast = nodes[len(nodes)-2]
		} else {
			secondToLast = nodes[0]
		}

		last.Next = first
		secondToLast.Next = nil
		first = last
		copy(nodes[1:], nodes[:len(nodes)-1])
		nodes[0] = last
	}

	return first
}

func sizeRecur(node *ListNode, acc int) int {
	if node == nil {
		return acc
	}
	return sizeRecur(node.Next, acc+1)
}

func size(head *ListNode) int {
	return sizeRecur(head, 0)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	nodes := nodesSeq(head)

	if k > len(nodes) {
		k %= len(nodes)
	}

	if len(nodes) == 1 || len(nodes) == k || k == 0 {
		return head
	}

	first := nodes[0]
	last := nodes[len(nodes)-1]
	prev := nodes[len(nodes)-1-k]
	next := nodes[len(nodes)-k]

	last.Next = first
	prev.Next = nil
	first = next

	return first
}
