package uniquebintrees

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	if tn == nil {
		return "x"
	}
	return fmt.Sprintf("%s<-%d->%s", tn.Left, tn.Val, tn.Right)
}

func helper(lo, hi int) []*TreeNode {
	var result []*TreeNode
	if lo > hi {
		result = append(result, nil)
		return result
	}
	for k := lo; k <= hi; k++ {
		leftSubTrees := helper(lo, k-1)
		rightSubTrees := helper(k+1, hi)
		for _, left := range leftSubTrees {
			for _, right := range rightSubTrees {
				node := &TreeNode{Val: k}
				node.Left = left
				node.Right = right
				result = append(result, node)
			}
		}
	}
	return result
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return helper(1, n)
}
