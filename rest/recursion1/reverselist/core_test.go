package reverselist

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

func TestReverseList(t *testing.T) {
	var tests = []struct {
		inputNodes []int
		wantNodes  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{nil, nil},
	}
	for _, test := range tests {
		list := makeList(test.inputNodes)
		want := makeList(test.wantNodes)
		if got := reverseList(list); !reflect.DeepEqual(got, want) {
			t.Errorf("reverseList(%v) = %v, want %v", test.inputNodes, got, want)
		}
	}
}
