package dupsubtrees

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tree *TreeNode) String() string {
	if tree == nil {
		return ""
	}
	return fmt.Sprintf("%s<-%d->%s", tree.Left, tree.Val, tree.Right)
}

func hash(m map[string][]*TreeNode, node *TreeNode) string {
	if node == nil {
		return ""
	}
	leftChildHash := hash(m, node.Left)
	rootHash := fmt.Sprintf("%d", node.Val)
	rightChildHash := hash(m, node.Right)
	var s strings.Builder
	if len(leftChildHash) > 0 {
		s.WriteString(fmt.Sprintf("%s<-", leftChildHash))
	}
	s.WriteString(rootHash)
	if len(rightChildHash) > 0 {
		s.WriteString(fmt.Sprintf("->%s", rightChildHash))
	}
	h := s.String()
	_, ok := m[h]
	if !ok {
		m[h] = []*TreeNode{node}
	} else {
		m[h] = append(m[h], node)
	}
	return h
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := make(map[string][]*TreeNode)
	m[hash(m, root)] = []*TreeNode{root}
	var result []*TreeNode
	for _, trees := range m {
		if len(trees) > 1 {
			result = append(result, trees[0])
		}
	}
	return result
}
