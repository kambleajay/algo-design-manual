package validate_binary_search_tree

import "testing"

func TestIsValidBST(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want bool
	}{
		{nil, true},
		{&TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}, true},
		{&TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 3}}, false},
		{&TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: -1}}, false},
		{&TreeNode{2, &TreeNode{Val: 2}, &TreeNode{Val: 2}}, false},
		{&TreeNode{5, &TreeNode{Val: 1}, &TreeNode{4, &TreeNode{Val: 3}, &TreeNode{Val: 6}}}, false},
		{&TreeNode{5, &TreeNode{Val: 4}, &TreeNode{6, &TreeNode{Val: 3}, &TreeNode{Val: 7}}}, false},
	}
	for _, test := range tests {
		if got := isValidBST(test.root); test.want != got {
			t.Errorf("isValidBST(%#v) should be %t, but was %t", test.root, test.want, got)
		}
	}
}
