package preorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalRecurHelper(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	preorderTraversalRecurHelper(root.Left, nums)
	preorderTraversalRecurHelper(root.Right, nums)
}

func preorderTraversalRecur(root *TreeNode) []int {
	nums := make([]int, 0)
	preorderTraversalRecurHelper(root, &nums)
	return nums
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

func preorderTraversalIter(root *TreeNode) []int {
	s := Stack{}
	var result []int
	if root == nil {
		return result
	}
	s.push(root)
	for !s.isEmpty() {
		x := s.pop()
		result = append(result, x.Val)
		if x.Right != nil {
			s.push(x.Right)
		}
		if x.Left != nil {
			s.push(x.Left)
		}
	}
	return result
}

func preorderTraversal(root *TreeNode) []int {
	return preorderTraversalIter(root)
}
