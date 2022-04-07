package construct_binary_tree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var tests = []struct {
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, &TreeNode{3, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}},
		{[]int{-1}, []int{-1}, &TreeNode{Val: -1}},
	}
	for _, test := range tests {
		if got := buildTree(test.preorder, test.inorder); !reflect.DeepEqual(test.want, got) {
			t.Errorf("buildTree(%v, %v) should be %v, but was %v", test.preorder, test.inorder, test.want, got)
		}
	}
}
