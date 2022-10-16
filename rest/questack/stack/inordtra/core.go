package inordtra

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	if n == nil {
		return "x"
	}
	return fmt.Sprintf("%s<-%d->%s", n.Left, n.Val, n.Right)
}

func inorderRecur(node *TreeNode, acc *[]int) {
	if node == nil {
		return
	}
	inorderRecur(node.Left, acc)
	*acc = append(*acc, node.Val)
	inorderRecur(node.Right, acc)
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorderRecur(root, &result)
	return result
}
