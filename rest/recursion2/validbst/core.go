package validbst

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	if tn == nil {
		return ""
	}
	return fmt.Sprintf("%s<-%d->%s", tn.Left, tn.Val, tn.Right)
}

func isBST(x *TreeNode, min, max *int) bool {
	if x == nil {
		return true
	}
	if min != nil && x.Val <= *min {
		return false
	}
	if max != nil && x.Val >= *max {
		return false
	}
	xval := new(int)
	*xval = x.Val
	return isBST(x.Left, min, xval) && isBST(x.Right, xval, max)
}

func isValidBST(root *TreeNode) bool {
	return isBST(root, nil, nil)
}
