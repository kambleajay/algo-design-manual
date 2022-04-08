package sort_list

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for head != nil {
		buf.WriteString(fmt.Sprintf("%d ", head.Val))
		head = head.Next
	}
	buf.WriteString("]")
	return buf.String()
}

type MinPQ struct {
	pq []int
}

func NewMinPQ() *MinPQ {
	pq := MinPQ{}
	pq.pq = []int{0}
	return &pq
}

func (minPQ *MinPQ) isEmpty() bool {
	return len(minPQ.pq) == 1
}

func (minPQ *MinPQ) swim(k int) {
	for k > 1 && minPQ.pq[k/2] > minPQ.pq[k] {
		minPQ.pq[k/2], minPQ.pq[k] = minPQ.pq[k], minPQ.pq[k/2]
		k /= 2
	}
}

func (minPQ *MinPQ) insert(x int) {
	minPQ.pq = append(minPQ.pq, x)
	minPQ.swim(len(minPQ.pq) - 1)
}

func (minPQ *MinPQ) sink(k int) {
	for 2*k <= len(minPQ.pq)-1 {
		j := 2 * k
		if j < len(minPQ.pq)-1 && minPQ.pq[j] > minPQ.pq[j+1] {
			j++
		}
		if minPQ.pq[k] < minPQ.pq[j] {
			break
		}
		minPQ.pq[k], minPQ.pq[j] = minPQ.pq[j], minPQ.pq[k]
		k = j
	}
}

func (minPQ *MinPQ) delMin() int {
	k := minPQ.pq[1]
	last := len(minPQ.pq) - 1
	minPQ.pq[1], minPQ.pq[last] = minPQ.pq[last], minPQ.pq[1]
	minPQ.pq = minPQ.pq[:last]
	minPQ.sink(1)
	return k
}

func sortList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pq := NewMinPQ()
	for head != nil {
		pq.insert(head.Val)
		head = head.Next
	}
	newHead := &ListNode{Val: pq.delMin()}
	node := newHead
	for !pq.isEmpty() {
		node.Next = &ListNode{Val: pq.delMin()}
		node = node.Next
	}
	return newHead
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}
	return head.Next
}

func getMid(head *ListNode) *ListNode {
	var slow *ListNode
	for head != nil && head.Next != nil {
		if slow == nil {
			slow = head
		} else {
			slow = slow.Next
		}
		head = head.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mid
}

func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)
	left := sortList2(head)
	right := sortList2(mid)
	return merge(left, right)
}

func sortList(head *ListNode) *ListNode {
	return sortList1(head)
}
