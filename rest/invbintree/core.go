package invbintree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) String() string {
	return fmt.Sprintf("%v<-%d->%v", root.Left, root.Val, root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Right)
	invertTree(root.Left)
	root.Left, root.Right = root.Right, root.Left
	return root
}
