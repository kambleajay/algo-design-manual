package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraphRecur(node *Node, seen map[int]*Node) *Node {
	earlierCopy, ok := seen[node.Val]
	if ok {
		return earlierCopy
	} else {
		newCopy := &Node{Val: node.Val}
		seen[node.Val] = newCopy
		for _, nb := range node.Neighbors {
			newCopy.Neighbors = append(newCopy.Neighbors, cloneGraphRecur(nb, seen))
		}
		return newCopy
	}
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	seen := make(map[int]*Node)
	return cloneGraphRecur(node, seen)
}
