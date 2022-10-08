package dupsubtrees

import (
	"reflect"
	"testing"
)

type Test struct {
	root *TreeNode
	want []*TreeNode
}

func tree1() *Test {
	one := &TreeNode{Val: 1}
	two1 := &TreeNode{Val: 2}
	four1 := &TreeNode{Val: 4}
	one.Left = two1
	two1.Left = four1
	three := &TreeNode{Val: 3}
	two2 := &TreeNode{Val: 2}
	four2 := &TreeNode{Val: 4}
	one.Right = three
	three.Left = two2
	two2.Left = four2
	four3 := &TreeNode{Val: 4}
	three.Right = four3
	return &Test{one, []*TreeNode{four1, two1}}
}

func tree2() *Test {
	two := &TreeNode{Val: 2}
	oneLeft := &TreeNode{Val: 1}
	oneRight := &TreeNode{Val: 1}
	two.Left = oneLeft
	two.Right = oneRight
	return &Test{two, []*TreeNode{oneLeft}}
}

func tree3() *Test {
	two1 := &TreeNode{Val: 2}
	two2 := &TreeNode{Val: 2}
	three1 := &TreeNode{Val: 3}
	two3 := &TreeNode{Val: 2}
	three2 := &TreeNode{Val: 3}
	two1.Left = two2
	two1.Right = two3
	two2.Left = three1
	two3.Left = three2
	return &Test{two1, []*TreeNode{three1, two2}}
}

func tree4() *Test {
	zero1 := &TreeNode{Val: 0}
	zero2 := &TreeNode{Val: 0}
	zero3 := &TreeNode{Val: 0}
	zero4 := &TreeNode{Val: 0}
	zero5 := &TreeNode{Val: 0}
	zero6 := &TreeNode{Val: 0}
	zero1.Left = zero2
	zero2.Left = zero3
	zero1.Right = zero4
	zero4.Right = zero5
	zero5.Right = zero6
	return &Test{zero1, []*TreeNode{zero3}}
}

func TestFindDuplicateSubtrees(t *testing.T) {
	var tests = []*Test{tree1(), tree2(), tree3(), tree4()}
	for _, test := range tests {
		if got := findDuplicateSubtrees(test.root); !reflect.DeepEqual(got, test.want) {
			t.Errorf("findDuplicateSubtrees(%v) = %v, want %v", test.root, got, test.want)
		}
	}
}
