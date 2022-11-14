package symmetric

import (
	"math"
	"testing"
)

func makeTree(nodes []int, k int) *TreeNode {
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

func TestIsSymmetric(t *testing.T) {
	var tests = []struct {
		nodes []int
		want  bool
	}{
		{[]int{0, 1, 2, 2, 3, 4, 4, 3}, true},
		{[]int{0, 1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5}, true},
		{[]int{0, 1, 2, 2, math.MaxInt64, 3, math.MaxInt64, 3}, false},
		{[]int{0, 9, -42, -42, math.MaxInt64, 76, 76, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 13, math.MaxInt64, 13, math.MaxInt64, math.MaxInt64}, false},
	}
	for _, test := range tests {
		root := makeTree(test.nodes, 1)
		if got := isSymmetric(root); got != test.want {
			t.Errorf("isSymmetric(%v) = %t, want %t", root, got, test.want)
		}
	}
}
