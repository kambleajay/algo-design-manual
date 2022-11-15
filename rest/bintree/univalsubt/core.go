package univalsubt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	items []*TreeNode
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(node *TreeNode) {
	q.items = append(q.items, node)
}

func (q *Queue) dequeue() *TreeNode {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func allNodesEqual(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}
	return allNodesEqual(node.Left, val) && allNodesEqual(node.Right, val)
}

func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count int
	var queue Queue
	queue.enqueue(root)
	for !queue.isEmpty() {
		node := queue.dequeue()
		if node.Left == nil && node.Right == nil {
			count++
		} else if allNodesEqual(node.Left, node.Val) && allNodesEqual(node.Right, node.Val) {
			count++
		}
		if node.Left != nil {
			queue.enqueue(node.Left)
		}
		if node.Right != nil {
			queue.enqueue(node.Right)
		}
	}
	return count
}
