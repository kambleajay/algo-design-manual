package middle_of_list

import "testing"

func makeList(xs []int) *ListNode {
	if len(xs) == 0 {
		return nil
	}
	head := &ListNode{Val: xs[0]}
	node := head
	for _, x := range xs[1:] {
		node.Next = &ListNode{Val: x}
		node = node.Next
	}
	return head
}

func TestMiddleNode(t *testing.T) {
	var tests = []struct {
		head *ListNode
		want int
	}{
		{makeList([]int{1, 2, 3, 4, 5}), 3},
		{makeList([]int{1, 2, 3, 4, 5, 6}), 4},
	}
	for _, test := range tests {
		if got := middleNode(test.head); test.want != got.Val {
			t.Errorf("middle should be equal %d, but was %d", test.want, got.Val)
		}
	}
}

func TestMiddleNodeNilList(t *testing.T) {
	if got := middleNode(nil); got != nil {
		t.Errorf("middleNode(nil) should be nil, but was %v", got)
	}
}
