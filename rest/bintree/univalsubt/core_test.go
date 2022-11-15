package univalsubt

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

func TestCountUnivalSubtrees(t *testing.T) {
	var tests = []struct {
		nodes []int
		want  int
	}{
		{[]int{0, 5, 1, 5, 5, 5, math.MaxInt64, 5}, 4},
		{nil, 0},
		{[]int{0, 5, 5, 5, 5, 5, math.MaxInt64, 5}, 6},
		{[]int{0, 1, math.MaxInt64, 2}, 1},
		{[]int{0, 1, 1, 1, 5, 5, math.MaxInt64, 5}, 3},
	}
	for _, test := range tests {
		root := makeTree(test.nodes, 1)
		if got := countUnivalSubtrees(root); got != test.want {
			t.Errorf("countUnivalSubtrees(%v) = %d, want %d", test.nodes, got, test.want)
		}
	}
}
