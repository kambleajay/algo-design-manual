package inscycle

type Node struct {
	Val  int
	Next *Node
}

func between(x int, prev *Node, curr *Node) bool {
	return prev.Val <= x && x <= curr.Val
}

func lessThanMin(x int, curr *Node) bool {
	return x <= curr.Val
}

func greaterThanMax(x int, prev *Node) bool {
	return x >= prev.Val
}

func isTailNode(prev, curr *Node) bool {
	return prev.Val > curr.Val
}

func canInsert(x int, prev, curr *Node) bool {
	atTailNode := isTailNode(prev, curr)
	return between(x, prev, curr) || atTailNode && lessThanMin(x, curr) || atTailNode && greaterThanMax(x, prev)
}

func insert(head *Node, x int) *Node {
	if head == nil {
		node := &Node{Val: x}
		node.Next = node
		return node
	}
	prev, curr := head, head.Next
	for ok := true; ok; ok = prev != head {
		if canInsert(x, prev, curr) {
			prev.Next = &Node{x, curr}
			return head
		}
		prev = curr
		curr = curr.Next
	}
	prev.Next = &Node{x, curr}
	return head
}
