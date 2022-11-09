package closebstval

import "testing"

func tree1() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n4.Left, n4.Right = n2, n5
	n2.Left, n2.Right = n1, n3
	return n4
}

func TestClosestValue(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		target float64
		want   int
	}{
		{tree1(), 3.714286, 4},
		{tree1(), 5.189, 5},
		{tree1(), 2.13, 2},
		{tree1(), 3.473, 3},
		{tree1(), 1.001, 1},
		{tree1(), 100.726, 5},
		{tree1(), -53.27, 1},
		{&TreeNode{Val: 1}, 4.428571, 1},
	}
	for _, test := range tests {
		if got := closestValue(test.tree, test.target); got != test.want {
			t.Errorf("closestValue(%v, %f) = %d, want %d", test.tree, test.target, got, test.want)
		}
	}
}
