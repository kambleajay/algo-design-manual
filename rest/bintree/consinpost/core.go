package consinpost

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeBuilder struct {
	inorder, postorder []int
	postOrderIndex     int
	valToIndex         map[int]int
}

func (tb *TreeBuilder) build(lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	val := tb.postorder[tb.postOrderIndex]
	rootInorderIndex := tb.valToIndex[val]
	node := &TreeNode{Val: val}
	tb.postOrderIndex--
	node.Right = tb.build(rootInorderIndex+1, hi)
	node.Left = tb.build(lo, rootInorderIndex-1)
	return node
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderValToIndex := make(map[int]int)
	for i, n := range inorder {
		inorderValToIndex[n] = i
	}
	tb := TreeBuilder{
		inorder:        inorder,
		postorder:      postorder,
		postOrderIndex: len(postorder) - 1,
		valToIndex:     inorderValToIndex,
	}
	return tb.build(0, len(postorder)-1)
}
