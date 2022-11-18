package swappairs

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

func TestSwapPairs(t *testing.T) {
	var tests = []struct {
		nodes []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{[]int{}, nil},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
		{[]int{1, 2}, []int{2, 1}},
	}
	for _, test := range tests {
		list := makeList(test.nodes)
		want := makeList(test.want)
		if got := swapPairs(list); !reflect.DeepEqual(got, want) {
			t.Errorf("swapPairs(%v) = %v, want %v", test.nodes, got, want)
		}
	}
}
