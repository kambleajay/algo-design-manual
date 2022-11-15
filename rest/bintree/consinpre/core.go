package consinpre

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeBuilder struct {
	preorder, inorder []int
	preOrderIndex     int
	valToIndex        map[int]int
}

func (tb *TreeBuilder) build(lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	val := tb.preorder[tb.preOrderIndex]
	idx := tb.valToIndex[val]
	node := &TreeNode{Val: val}
	tb.preOrderIndex++
	node.Left = tb.build(lo, idx-1)
	node.Right = tb.build(idx+1, hi)
	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	valToIdx := make(map[int]int)
	for i, v := range inorder {
		valToIdx[v] = i
	}
	tb := &TreeBuilder{
		preorder:      preorder,
		inorder:       inorder,
		preOrderIndex: 0,
		valToIndex:    valToIdx,
	}
	return tb.build(0, len(preorder)-1)
}
