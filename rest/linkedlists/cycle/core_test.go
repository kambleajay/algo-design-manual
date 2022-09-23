package cycle

import "testing"

func list1() *ListNode {
	n4 := &ListNode{Val: -4}
	n3 := &ListNode{Val: 0, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 3, Next: n2}
	n4.Next = n2
	return n1
}

func list2() *ListNode {
	n2 := &ListNode{Val: 2}
	n1 := &ListNode{Val: 1, Next: n2}
	n2.Next = n1
	return n1
}

func list3() *ListNode {
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	return n1
}

func TestHasCycle(t *testing.T) {
	var tests = []struct {
		list *ListNode
		want bool
	}{
		{list1(), true},
		{list2(), true},
		{&ListNode{Val: 1}, false},
		{list3(), false},
	}
	for _, test := range tests {
		if got := hasCycle(test.list); got != test.want {
			t.Errorf("hasCycle(%v) = %t, want %t", test.list, got, test.want)
		}
	}
}
