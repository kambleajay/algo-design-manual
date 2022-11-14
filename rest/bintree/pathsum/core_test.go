package pathsum

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

func TestHasPathSum(t *testing.T) {
	var tests = []struct {
		nodes     []int
		targetSum int
		want      bool
	}{
		{[]int{0, 5, 4, 8, 11, math.MaxInt64, 13, 4, 7, 2, math.MaxInt64, math.MaxInt64, math.MaxInt64, 1}, 22, true},
		{[]int{0, 1, 2, 3}, 5, false},
		{[]int{0, 1, math.MaxInt64, 2}, 1, false},
		{nil, 0, false},
	}
	for _, test := range tests {
		root := makeTree(test.nodes, 1)
		if got := hasPathSum(root, test.targetSum); got != test.want {
			t.Errorf("hasPathSum(%v, %d) = %t, want %t", test.nodes, test.targetSum, got, test.want)
		}
	}
}
