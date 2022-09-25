package cycle

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

func TestHasCycle(t *testing.T) {
	var tests = []struct {
		list *ll.ListNode
		want bool
	}{
		{list1(), true},
		{list2(), true},
		{&ll.ListNode{Val: 1}, false},
		{list3(), false},
	}
	for _, test := range tests {
		if got := hasCycle(test.list); got != test.want {
			t.Errorf("hasCycle(%v) = %t, want %t", test.list, got, test.want)
		}
	}
}
