package invbintree

import (
	"reflect"
	"testing"
)

func emptyNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func node(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func TestInvertTree(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want *TreeNode
	}{
		{node(2, emptyNode(1), emptyNode(3)), node(2, emptyNode(3), emptyNode(1))},
		{node(4, node(2, emptyNode(1), emptyNode(3)), node(7, emptyNode(6), emptyNode(9))),
			node(4, node(7, emptyNode(9), emptyNode(6)), node(2, emptyNode(3), emptyNode(1)))},
	}
	for _, test := range tests {
		if got := invertTree(test.root); !reflect.DeepEqual(got, test.want) {
			t.Errorf("invertTree(%v) = %v, want %v", test.root, got, test.want)
		}
	}
}
