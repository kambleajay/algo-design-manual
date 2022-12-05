package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	p, q *TreeNode
}

type Queue struct {
	items []*Pair
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(pair *Pair) {
	q.items = append(q.items, pair)
}

func (q *Queue) dequeue() *Pair {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func isValid(x, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if x.Val != y.Val {
		return false
	}
	return true
}

func isSameTreeIter(p *TreeNode, q *TreeNode) bool {
	var queue Queue
	queue.enqueue(&Pair{p, q})
	for !queue.isEmpty() {
		pair := queue.dequeue()
		x, y := pair.p, pair.q
		if !isValid(x, y) {
			return false
		}
		if x != nil && y != nil {
			queue.enqueue(&Pair{x.Left, y.Left})
			queue.enqueue(&Pair{x.Right, y.Right})
		}
	}
	return true
}

func isSameTreeRecur(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeIter(p, q)
}
