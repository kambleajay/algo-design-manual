package copyrandlist

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%d (%d) -> %s", n.Val, n.Random.Val, n.Next)
}

func copyNode(node *Node, seen map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	var result *Node
	cp, ok := seen[node]
	if ok {
		result = cp
	} else {
		newNode := &Node{Val: node.Val}
		seen[node] = newNode
		result = newNode
	}
	return result
}

func copyRandomList(head *Node) *Node {
	seen := make(map[*Node]*Node)
	if head == nil {
		return nil
	}
	tail := head
	node := &Node{Val: tail.Val}
	seen[tail] = node
	for tail != nil {
		node.Random = copyNode(tail.Random, seen)
		node.Next = copyNode(tail.Next, seen)
		tail = tail.Next
		node = node.Next
	}
	return seen[head]
}
