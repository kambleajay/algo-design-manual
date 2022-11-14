package postorder

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

func TestPostorderTraversal(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want []int
	}{
		{tree1(), []int{3, 2, 1}},
		{nil, nil},
		{&TreeNode{Val: 1}, []int{1}},
	}
	for _, test := range tests {
		if got := postorderTraversal(test.root); !reflect.DeepEqual(got, test.want) {
			t.Errorf("postorderTraversal(%v) = %v, want %v", test.root, got, test.want)
		}
	}
}
