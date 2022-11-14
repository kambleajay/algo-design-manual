package postorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalRecur(x *TreeNode, result *[]int) {
	if x == nil {
		return
	}
	postorderTraversalRecur(x.Left, result)
	postorderTraversalRecur(x.Right, result)
	*result = append(*result, x.Val)
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	postorderTraversalRecur(root, &result)
	return result
}
