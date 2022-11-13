package inorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	items []*TreeNode
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) push(node *TreeNode) {
	s.items = append(s.items, node)
}

func (s *Stack) pop() *TreeNode {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func inorderTraversalIter(root *TreeNode) []int {
	var stack Stack
	var result []int
	x := root
	for x != nil || !stack.isEmpty() {
		for x != nil {
			stack.push(x)
			x = x.Left
		}
		x = stack.pop()
		result = append(result, x.Val)
		x = x.Right
	}
	return result
}

func inorderTraversal(root *TreeNode) []int {
	return inorderTraversalIter(root)
}
