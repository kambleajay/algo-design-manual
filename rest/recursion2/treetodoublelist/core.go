package treetodoublelist

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Stack struct {
	items []*Node
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) push(item *Node) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() *Node {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func inorder(root *Node) []*Node {
	var stack Stack
	var result []*Node
	x := root
	for x != nil || !stack.isEmpty() {
		for x != nil {
			stack.push(x)
			x = x.Left
		}
		x = stack.pop()
		result = append(result, x)
		x = x.Right
	}
	return result
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	iord := inorder(root)
	for i := 1; i < len(iord); i++ {
		x, y := iord[i-1], iord[i]
		x.Right = y
		y.Left = x
	}
	iord[0].Left = iord[len(iord)-1]
	iord[len(iord)-1].Right = iord[0]
	return iord[0]
}
