package lowcommonancestor

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	if tn == nil {
		return ""
	}
	return fmt.Sprintf("%s<-%d->%s", tn.Left, tn.Val, tn.Right)
}

func adj(node *TreeNode) []*TreeNode {
	var result []*TreeNode
	if node.Left != nil {
		result = append(result, node.Left)
	}
	if node.Right != nil {
		result = append(result, node.Right)
	}
	return result
}

func dfs(from *TreeNode, seen map[int]bool, edgeTo map[*TreeNode]*TreeNode) {
	if from == nil {
		return
	}
	seen[from.Val] = true
	for _, to := range adj(from) {
		if !seen[to.Val] {
			dfs(to, seen, edgeTo)
			edgeTo[to] = from
		}
	}
}

func pathTo(v *TreeNode, root *TreeNode, edgeTo map[*TreeNode]*TreeNode) []*TreeNode {
	var result []*TreeNode
	for x := v; x != root; x = edgeTo[x] {
		result = append(result, x)
	}
	result = append(result, root)
	return result
}

func contains(a []*TreeNode, x *TreeNode) bool {
	for _, v := range a {
		if x == v {
			return true
		}
	}
	return false
}

func lowestCommon(x []*TreeNode, y []*TreeNode) *TreeNode {
	if len(x) > len(y) {
		return lowestCommon(y, x)
	}
	seenInX := make(map[*TreeNode]bool)
	for _, e := range x {
		seenInX[e] = true
	}
	for _, e := range y {
		if seenInX[e] {
			return e
		}
	}
	return nil
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	seen := make(map[int]bool)
	edgeTo := make(map[*TreeNode]*TreeNode)
	dfs(root, seen, edgeTo)
	pathToP := pathTo(p, root, edgeTo)
	pathToQ := pathTo(q, root, edgeTo)
	if contains(pathToQ, p) {
		return p
	} else if contains(pathToP, q) {
		return q
	}
	return lowestCommon(pathToP, pathToQ)
}

type Stack struct {
	items []*TreeNode
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) push(node *TreeNode) {
	s.items = append(s.items, node)
}

func (s *Stack) pop() *TreeNode {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func notIn(m map[*TreeNode]*TreeNode, node *TreeNode) bool {
	_, ok := m[node]
	return !ok
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var stack Stack
	parents := make(map[*TreeNode]*TreeNode)
	parents[root] = nil
	stack.push(root)
	for notIn(parents, p) || notIn(parents, q) {
		node := stack.pop()
		if node.Left != nil {
			parents[node.Left] = node
			stack.push(node.Left)
		}
		if node.Right != nil {
			parents[node.Right] = node
			stack.push(node.Right)
		}
	}
	pathToP := make(map[*TreeNode]bool)
	for x := p; x != root; x = parents[x] {
		pathToP[x] = true
	}
	pathToP[root] = true
	var answer *TreeNode
	for answer = q; !pathToP[answer]; answer = parents[answer] {
	}
	return answer
}
