package validate_binary_search_tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n TreeNode) String() string {
	return fmt.Sprintf("%v <- %d -> %v", n.Left, n.Val, n.Right)
}

func isValidBST(root *TreeNode) bool {
	var prev *int
	var inorder func(root *TreeNode) bool
	inorder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !inorder(root.Left) {
			return false
		}
		if prev != nil && root.Val <= *prev {
			return false
		}
		prev = &root.Val
		return inorder(root.Right)
	}
	return inorder(root)
}
