package intersection

import (
	ll "algo/rest/linkedlists"
	"testing"
)

type Test struct {
	headA *ll.ListNode
	headB *ll.ListNode
	want  *ll.ListNode
}

func list1() *Test {
	c3 := &ll.ListNode{Val: 5}
	c2 := &ll.ListNode{4, c3}
	c1 := &ll.ListNode{8, c2}
	a2 := &ll.ListNode{1, c1}
	a1 := &ll.ListNode{4, a2}
	b3 := &ll.ListNode{1, c1}
	b2 := &ll.ListNode{6, b3}
	b1 := &ll.ListNode{5, b2}
	return &Test{a1, b1, c1}
}

func list2() *Test {
	c2 := &ll.ListNode{Val: 4}
	c1 := &ll.ListNode{2, c2}
	a3 := &ll.ListNode{1, c1}
	a2 := &ll.ListNode{9, a3}
	a1 := &ll.ListNode{1, a2}
	b1 := &ll.ListNode{3, c1}
	return &Test{a1, b1, c1}
}

func list3() *Test {
	a3 := &ll.ListNode{Val: 4}
	a2 := &ll.ListNode{6, a3}
	a1 := &ll.ListNode{2, a2}
	b2 := &ll.ListNode{Val: 5}
	b1 := &ll.ListNode{1, b2}
	return &Test{a1, b1, nil}
}

func TestGetIntersectionNode(t *testing.T) {
	var tests = []*Test{list1(), list2(), list3()}
	for _, test := range tests {
		if got := getIntersectionNode(test.headA, test.headB); got != test.want {
			t.Errorf("getIntersectionNode(%v, %v) = %v, want %v", test.headA, test.headB, got, test.want)
		}
	}
}
