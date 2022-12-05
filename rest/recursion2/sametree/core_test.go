package sametree

import (
	"math"
	"testing"
)

func makeTree(nodes []int, k int) *TreeNode {
	if nodes == nil {
		return nil
	}
	if nodes[k] == math.MaxInt64 {
		return nil
	}
	node := &TreeNode{Val: nodes[k]}
	left, right := 2*k, (2*k)+1
	if left < len(nodes) {
		node.Left = makeTree(nodes, left)
	}
	if right < len(nodes) {
		node.Right = makeTree(nodes, right)
	}
	return node
}

func TestSameTree(t *testing.T) {
	var tests = []struct {
		pNodes, qNodes []int
		want           bool
	}{
		{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}, true},
		{[]int{0, 1, 2, math.MaxInt64}, []int{0, 1, math.MaxInt64, 2}, false},
		{[]int{0, 1, 2, 1}, []int{0, 1, 1, 2}, false},
		{[]int{0, 1}, []int{0, 1}, true},
		{nil, nil, true},
	}
	for _, test := range tests {
		p := makeTree(test.pNodes, 1)
		q := makeTree(test.qNodes, 1)
		if got := isSameTree(p, q); got != test.want {
			t.Errorf("isSameTree(%v, %v) = %t, want %t", test.pNodes, test.qNodes, got, test.want)
		}
	}
}
