package rotate_list

import (
	"reflect"
	"testing"
)

func values(head *ListNode) []int {
	var vals []int
	x := head
	for x != nil {
		vals = append(vals, x.Val)
		x = x.Next
	}
	return vals
}

func makeListRecur(node *ListNode, values []int) {
	if len(values) > 0 {
		x := &ListNode{Val: values[0]}
		node.Next = x
		makeListRecur(x, values[1:])
	}
	return
}

func makeList(values []int) *ListNode {
	head := &ListNode{Val: values[0]}
	makeListRecur(head, values[1:])
	return head
}

func TestRotateRight(t *testing.T) {
	var tests = []struct {
		list *ListNode
		k    int
		want []int
	}{
		{nil, 1, nil},
		{makeList([]int{1, 2, 3}), 0, []int{1, 2, 3}},
		{makeList([]int{1}), 0, []int{1}},
		{makeList([]int{1}), 3, []int{1}},
		{makeList([]int{1, 2}), 1, []int{2, 1}},
		{makeList([]int{1, 2}), 2, []int{1, 2}},
		{makeList([]int{1, 2, 3, 4, 5}), 1, []int{5, 1, 2, 3, 4}},
		{makeList([]int{1, 2, 3, 4, 5}), 2, []int{4, 5, 1, 2, 3}},
		{makeList([]int{1, 2, 3, 4, 5}), 3, []int{3, 4, 5, 1, 2}},
		{makeList([]int{1, 2, 3, 4, 5}), 4, []int{2, 3, 4, 5, 1}},
		{makeList([]int{1, 2, 3, 4, 5}), 5, []int{1, 2, 3, 4, 5}},
		{makeList([]int{1, 2, 3, 4, 5}), 6, []int{5, 1, 2, 3, 4}},
		{makeList([]int{1, 2, 3, 4, 5}), 7, []int{4, 5, 1, 2, 3}},
		{makeList([]int{1, 2, 3, 4, 5}), 8, []int{3, 4, 5, 1, 2}},
		{makeList([]int{1, 2, 3, 4, 5}), 9, []int{2, 3, 4, 5, 1}},
		{makeList([]int{1, 2, 3, 4, 5}), 10, []int{1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		//fmt.Printf("%v, %d\n", test.list, test.k)
		got := rotateRight(test.list, test.k)
		vals := values(got)
		if !reflect.DeepEqual(vals, test.want) {
			t.Errorf("for %v rotated list must be %v, but was %v", test.list, test.want, vals)
		}
	}
}
