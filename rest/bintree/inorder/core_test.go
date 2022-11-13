package inorder

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

func TestInorderTraversal(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want []int
	}{
		{tree1(), []int{1, 3, 2}},
		{nil, nil},
		{&TreeNode{Val: 1}, []int{1}},
	}
	for _, test := range tests {
		if got := inorderTraversal(test.root); !reflect.DeepEqual(got, test.want) {
			t.Errorf("inorderTraversal(%v) = %v, want %v", test.root, got, test.want)
		}
	}
}
