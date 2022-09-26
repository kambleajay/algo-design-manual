package flatten

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func (node *Node) String() string {
	if node == nil {
		return ""
	}
	var s string
	if node.Child != nil {
		s = fmt.Sprintf("%d (%v) -> %v", node.Val, node.Child, node.Next)
	} else {
		s = fmt.Sprintf("%d -> %v", node.Val, node.Next)
	}
	return s
}

func flattenRecur(pred *Node, curr *Node) *Node {
	if curr == nil {
		return pred
	}
	curr.Prev = pred
	pred.Next = curr
	succ := curr.Next

	tail := flattenRecur(curr, curr.Child)
	curr.Child = nil

	return flattenRecur(tail, succ)
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	head := &Node{}
	flattenRecur(head, root)
	head.Next.Prev = nil
	return head.Next
}
