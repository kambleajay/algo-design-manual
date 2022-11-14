package maxdepth

import "testing"

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
	n1.Right = n2
	return n1
}

func tree3() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Left = n2
	n2.Left = n3
	return n1
}

func tree4() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Right = n2
	n2.Right = n3
	return n1
}

func tree5() *TreeNode {
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

func TestMaxDepth(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want int
	}{
		{tree1(), 3},
		{tree2(), 2},
		{tree3(), 3},
		{tree4(), 3},
		{tree5(), 3},
		{&TreeNode{Val: 1}, 1},
		{nil, 0},
	}
	for _, test := range tests {
		if got := maxDepth(test.root); got != test.want {
			t.Errorf("maxDepth(%v) = %d, want %d", test.root, got, test.want)
		}
	}
}
