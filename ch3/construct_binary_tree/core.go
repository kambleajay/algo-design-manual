package construct_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var preOrderIndex int
	inOrderKeyToIndex := make(map[int]int)
	var arrayToTree func(int, int) *TreeNode

	arrayToTree = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		rootKey := preorder[preOrderIndex]
		preOrderIndex++
		root := TreeNode{Val: rootKey}
		rootIndexInInorder := inOrderKeyToIndex[rootKey]
		root.Left = arrayToTree(left, rootIndexInInorder-1)
		root.Right = arrayToTree(rootIndexInInorder+1, right)
		return &root
	}

	for i, elem := range inorder {
		inOrderKeyToIndex[elem] = i
	}

	return arrayToTree(0, len(preorder)-1)
}
