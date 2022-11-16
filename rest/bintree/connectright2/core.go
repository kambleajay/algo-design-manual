package connectright2

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%s<-%d(%s)->%s", n.Left, n.Val, n.Next, n.Right)
}

type Connector struct {
	leftmost      *Node
	nextLevelCurr *Node
}

func (c *Connector) processChild(node *Node) {
	if node != nil {
		if c.nextLevelCurr == nil {
			c.leftmost = node
		} else {
			c.nextLevelCurr.Next = node
		}
		c.nextLevelCurr = node
	}
}

func (c *Connector) connectNexts(root *Node) *Node {
	if root == nil {
		return nil
	}
	c.leftmost = root
	curr := c.leftmost
	for c.leftmost != nil {
		curr, c.leftmost = c.leftmost, nil
		c.nextLevelCurr = nil
		for curr != nil {
			c.processChild(curr.Left)
			c.processChild(curr.Right)
			curr = curr.Next
		}
	}
	return root
}

func connect(root *Node) *Node {
	var c Connector
	return c.connectNexts(root)
}
