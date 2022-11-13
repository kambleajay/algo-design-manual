package preorder

import (
	"reflect"
	"testing"
)

func tree1() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Right = n2
	n2.Left = n3
	return n1
}

func tree2() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n1.Left = n4
	n4.Left = n2
	n1.Right = n3
	return n1
}

func TestPreOrderTraversal(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want []int
	}{
		{tree1(), []int{1, 2, 3}},
		{tree2(), []int{1, 4, 2, 3}},
		{nil, nil},
		{&TreeNode{Val: 1}, []int{1}},
	}
	for _, test := range tests {
		if got := preorderTraversal(test.root); !reflect.DeepEqual(got, test.want) {
			t.Errorf("preorderTraversal(%v) = %v, want %v", test.root, got, test.want)
		}
	}
}
