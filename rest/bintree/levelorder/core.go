package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	items [][]*TreeNode
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(node []*TreeNode) {
	q.items = append(q.items, node)
}

func (q *Queue) dequeue() []*TreeNode {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func adj(nodes []*TreeNode) []*TreeNode {
	var result []*TreeNode
	for _, node := range nodes {
		if node.Left != nil {
			result = append(result, node.Left)
		}
		if node.Right != nil {
			result = append(result, node.Right)
		}
	}
	return result
}

func values(nodes []*TreeNode) []int {
	var result []int
	for _, node := range nodes {
		result = append(result, node.Val)
	}
	return result
}

func levelOrder(root *TreeNode) [][]int {
	var queue Queue
	var result [][]int
	if root == nil {
		return result
	}
	queue.enqueue([]*TreeNode{root})
	result = append(result, []int{root.Val})
	for !queue.isEmpty() {
		from := queue.dequeue()
		nextLevelNodes := adj(from)
		if len(nextLevelNodes) > 0 {
			result = append(result, values(nextLevelNodes))
			queue.enqueue(nextLevelNodes)
		}
	}
	return result
}
