package mergelists

import (
	"reflect"
	"testing"
)

func makeList(nodes []int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}
	node := &ListNode{Val: nodes[0]}
	node.Next = makeList(nodes[1:])
	return node
}

func TestMergeTwoLists(t *testing.T) {
	var tests = []struct {
		list1Nodes []int
		list2Nodes []int
		wantNodes  []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{nil, nil, nil},
		{nil, []int{0}, []int{0}},
		{[]int{0}, nil, []int{0}},
	}
	for _, test := range tests {
		list1 := makeList(test.list1Nodes)
		list2 := makeList(test.list2Nodes)
		want := makeList(test.wantNodes)
		if got := mergeTwoLists(list1, list2); !reflect.DeepEqual(got, want) {
			t.Errorf("mergeTwoLists(%v, %v) = %v, want %v", list1, list2, got, want)
		}
	}
}
