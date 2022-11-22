package validbst

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

func TestIsValidBST(t *testing.T) {
	var tests = []struct {
		inputNodes []int
		want       bool
	}{
		{[]int{0, 2, 1, 3}, true},
		{nil, true},
		{[]int{0, 5, 1, 4, math.MaxInt64, math.MaxInt64, 3, 6}, false},
		{[]int{0, 2, 2, 2}, false},
	}
	for _, test := range tests {
		input := makeTree(test.inputNodes, 1)
		if got := isValidBST(input); got != test.want {
			t.Errorf("isValidBST(%v) = %t, want %t", input, got, test.want)
		}
	}
}
