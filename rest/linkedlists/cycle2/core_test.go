package cycle2

import (
	ll "algo/rest/linkedlists"
	"testing"
)

func list1() *ll.ListNode {
	n4 := &ll.ListNode{Val: -4}
	n3 := &ll.ListNode{Val: 0, Next: n4}
	n2 := &ll.ListNode{Val: 2, Next: n3}
	n1 := &ll.ListNode{Val: 3, Next: n2}
	n4.Next = n2
	return n1
}

func list2() *ll.ListNode {
	n2 := &ll.ListNode{Val: 2}
	n1 := &ll.ListNode{Val: 1, Next: n2}
	n2.Next = n1
	return n1
}

func list3() *ll.ListNode {
	n3 := &ll.ListNode{Val: 3}
	n2 := &ll.ListNode{Val: 2, Next: n3}
	n1 := &ll.ListNode{Val: 1, Next: n2}
	return n1
}

func TestDetectCycle(t *testing.T) {
	var tests = []struct {
		list *ll.ListNode
		want int
	}{
		{list1(), 2},
		{list2(), 1},
	}
	for _, test := range tests {
		if got := detectCycle(test.list); got.Val != test.want {
			t.Errorf("detectCycle(%v) = %d, want %d", test.list, got.Val, test.want)
		}
	}
}

func TestDetectCycleNoCycle(t *testing.T) {
	var tests = []struct {
		list *ll.ListNode
	}{
		{list3()},
		{&ll.ListNode{Val: 1}},
	}
	for _, test := range tests {
		if got := detectCycle(test.list); got != nil {
			t.Errorf("detectCycle(%v) = %v, want %v", test.list, got, nil)
		}
	}
}
