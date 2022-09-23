package cycle2

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

func TestDetectCycle(t *testing.T) {
	var tests = []struct {
		list *ListNode
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
		list *ListNode
	}{
		{list3()},
		{&ListNode{Val: 1}},
	}
	for _, test := range tests {
		if got := detectCycle(test.list); got != nil {
			t.Errorf("detectCycle(%v) = %v, want %v", test.list, got, nil)
		}
	}
}
