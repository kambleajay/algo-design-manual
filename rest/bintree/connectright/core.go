package connectright

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connectHelper(nodes []*Node) {
	if len(nodes) == 0 {
		return
	}
	var children []*Node
	if len(nodes) > 1 {
		for i := 0; i < len(nodes)-1; i++ {
			nodes[i].Next = nodes[i+1]
		}
	}
	for _, node := range nodes {
		if node.Left != nil {
			children = append(children, node.Left)
		}
		if node.Right != nil {
			children = append(children, node.Right)
		}
	}
	connectHelper(children)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	leftmost := root
	for leftmost.Left != nil {
		head := leftmost
		for head != nil {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
			head = head.Next
		}
		leftmost = leftmost.Left
	}
	return root
}
