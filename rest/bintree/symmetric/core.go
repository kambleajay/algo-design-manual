package symmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bothEmpty(left, right *TreeNode) bool {
	return left == nil && right == nil
}

func oneEmpty(left, right *TreeNode) bool {
	return (left == nil && right != nil) || (left != nil && right == nil)
}

func mirrored(left, right *TreeNode) bool {
	if bothEmpty(left, right) {
		return true
	}
	if oneEmpty(left, right) {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return mirrored(left.Left, right.Right) && mirrored(left.Right, right.Left)
}

func isSymmetricRecur(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirrored(root.Left, root.Right)
}

type Pair struct {
	x, y *TreeNode
}

type Queue struct {
	items []*Pair
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(p *Pair) {
	q.items = append(q.items, p)
}

func (q *Queue) dequeue() *Pair {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func isSymmetricIter(root *TreeNode) bool {
	var queue Queue
	if root == nil {
		return true
	}
	queue.enqueue(&Pair{root.Left, root.Right})
	for !queue.isEmpty() {
		p := queue.dequeue()
		if bothEmpty(p.x, p.y) {
			continue
		}
		if oneEmpty(p.x, p.y) {
			return false
		}
		if p.x.Val != p.y.Val {
			return false
		}
		queue.enqueue(&Pair{p.x.Left, p.y.Right})
		queue.enqueue(&Pair{p.x.Right, p.y.Left})
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricIter(root)
}
