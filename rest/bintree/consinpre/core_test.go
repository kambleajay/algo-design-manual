package consinpre

import (
	"math"
	"reflect"
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

func TestBuildTree(t *testing.T) {
	var tests = []struct {
		preorder, inorder, nodes []int
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []int{0, 3, 9, 20, math.MaxInt64, math.MaxInt64, 15, 7}},
		{[]int{-1}, []int{-1}, []int{0, -1}},
	}
	for _, test := range tests {
		want := makeTree(test.nodes, 1)
		if got := buildTree(test.preorder, test.inorder); !reflect.DeepEqual(got, want) {
			t.Errorf("buildTree(%v, %v) = %v, want %v", test.preorder, test.inorder, got, want)
		}
	}
}
