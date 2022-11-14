package maxdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftBranchDepth := maxDepth(root.Left)
	rightBranchDepth := maxDepth(root.Right)
	return max(leftBranchDepth, rightBranchDepth) + 1
}
