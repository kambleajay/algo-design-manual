package uniquebintrees

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

func containsAll(xs []*TreeNode, ys []*TreeNode) bool {
	m := make(map[TreeNode]bool)
	for _, x := range xs {
		m[*x] = true
	}
	for _, y := range ys {
		if !m[*y] {
			return false
		}
	}
	return true
}

func TestGenerateTrees(t *testing.T) {
	var tests = []struct {
		n         int
		wantNodes [][]int
	}{
		{1, [][]int{{0, 1}}},
		{2, [][]int{{0, 1, math.MaxInt64, 2}, {0, 2, 1, math.MaxInt64}}},
		{3, [][]int{{0, 1, math.MaxInt64, 3, math.MaxInt64, math.MaxInt64, 2, math.MaxInt64},
			{0, 1, math.MaxInt64, 2, math.MaxInt64, math.MaxInt64, math.MaxInt64, 3},
			{0, 2, 1, 3},
			{0, 3, 2, math.MaxInt64, 1, math.MaxInt64},
			{0, 3, 1, math.MaxInt64, math.MaxInt64, 2}}},
	}
	for _, test := range tests {
		var want []*TreeNode
		for _, wn := range test.wantNodes {
			want = append(want, makeTree(wn, 1))
		}
		if got := generateTrees(test.n); !containsAll(got, want) {
			t.Errorf("generateTrees(%d) = %v, want %v", test.n, got, want)
		}
	}
}
