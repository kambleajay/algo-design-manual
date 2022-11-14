package symmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bothEmpty(left, right *TreeNode) bool {
	return left == nil && right == nil
}

func oneEmpty(left, right *TreeNode) bool {
	return (left == nil && right != nil) || (left != nil && right == nil)
}

func mirrored(left, right *TreeNode) bool {
	if bothEmpty(left, right) {
		return true
	}
	if oneEmpty(left, right) {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return mirrored(left.Left, right.Right) && mirrored(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirrored(root.Left, root.Right)
}
