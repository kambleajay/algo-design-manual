package levelorder

import (
	"reflect"
	"testing"
)

func tree1() *TreeNode {
	n3 := &TreeNode{Val: 3}
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}
	n3.Left, n3.Right = n9, n20
	n20.Left, n20.Right = n15, n7
	return n3
}

func tree2() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Left = n2
	n2.Left = n3
	return n1
}

func tree3() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Right = n2
	n2.Right = n3
	return n1
}

func tree4() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Left, n3.Right = n6, n7
	return n1
}

func TestLevelOrder(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want [][]int
	}{
		{tree1(), [][]int{{3}, {9, 20}, {15, 7}}},
		{&TreeNode{Val: 1}, [][]int{{1}}},
		{nil, nil},
		{tree2(), [][]int{{1}, {2}, {3}}},
		{tree3(), [][]int{{1}, {2}, {3}}},
		{tree4(), [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
	}
	for _, test := range tests {
		if got := levelOrder(test.root); !reflect.DeepEqual(got, test.want) {
			t.Errorf("levelOrder(%v) = %v, want %v", test.root, got, test.want)
		}
	}
}
