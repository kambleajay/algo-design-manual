package backtracking

type Node struct {
	word     string
	children map[byte]*Node
}

func NewNode() *Node {
	children := make(map[byte]*Node)
	return &Node{children: children}
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	node := NewNode()
	return &Trie{root: node}
}

func (tri *Trie) put(key string) {
	x := tri.root
	for _, ch := range []byte(key) {
		child, ok := x.children[ch]
		if !ok {
			node := NewNode()
			x.children[ch] = node
			x = node
		} else {
			x = child
		}
	}
	x.word = key
}

func (tri *Trie) get(key string) *Node {
	x := tri.root
	for _, ch := range []byte(key) {
		child, ok := x.children[ch]
		if !ok {
			return nil
		} else {
			x = child
		}
	}
	return x
}

func (tri *Trie) contains(key string) bool {
	x := tri.get(key)
	result := x != nil && len(x.word) > 0
	if result && len(x.children) == 0 {
		last := key[len(key)-1]
		delete(x.children, last)
	}
	return result
}

func (tri *Trie) containsPrefix(key string) bool {
	x := tri.get(key)
	return x != nil
}

type Cell struct {
	x, y int
}

type Solution struct {
	trie   *Trie
	board  [][]byte
	words  []string
	result []string
	a      []byte
	moves  []*Cell
	m, n   int
}

func NewSolution(board [][]byte, words []string) Solution {
	trie := NewTrie()
	for _, word := range words {
		trie.put(word)
	}
	a := make([]byte, 10)
	m, n := len(board), len(board[0])
	moves := make([]*Cell, (m*n)+1)
	return Solution{trie: trie, board: board, words: words, a: a, moves: moves, m: m, n: n}
}

func (s *Solution) resize(capacity int) {
	tmp := make([]byte, capacity)
	copy(tmp, s.a)
	s.a = tmp
}

func (s *Solution) isSolution(node *Node) bool {
	return len(node.word) > 0
}

func (s *Solution) processSolution(node *Node) {
	s.result = append(s.result, node.word)
	node.word = ""
}

func (s *Solution) allValidCells() []*Cell {
	var cells []*Cell
	for i := 0; i < s.m; i++ {
		for j := 0; j < s.n; j++ {
			ch := s.board[i][j]
			if s.trie.containsPrefix(string(ch)) {
				cells = append(cells, &Cell{i, j})
			}
		}
	}
	return cells
}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (s *Solution) isValidRow(r int) bool {
	return r >= 0 && r < s.m
}

func (s *Solution) isValidCol(c int) bool {
	return c >= 0 && c < s.n
}

func (s *Solution) isValidPrefix(k, x, y int, node *Node) bool {
	ch := s.board[x][y]
	_, ok := node.children[ch]
	return ok
}

func (s *Solution) cellUsed(k, x, y int) bool {
	cell := Cell{x, y}
	for i := 1; i < k; i++ {
		if *s.moves[i] == cell {
			return true
		}
	}
	return false
}

func (s *Solution) constructCandidates(k int, currCell *Cell, node *Node) []*Cell {
	if k == 1 {
		return s.allValidCells()
	}
	var candidates []*Cell
	for _, dir := range dirs {
		x, y := dir[0], dir[1]
		px, py := currCell.x+x, currCell.y+y
		if s.isValidRow(px) && s.isValidCol(py) && !s.cellUsed(k, px, py) && s.isValidPrefix(k, px, py, node) {
			candidates = append(candidates, &Cell{px, py})
		}
	}
	return candidates
}

func (s *Solution) backtrack(k int, currCell *Cell, node *Node) {
	if s.isSolution(node) {
		s.processSolution(node)
	}
	k++
	for _, candidate := range s.constructCandidates(k, currCell, node) {
		if k == len(s.a) {
			s.resize(2 * len(s.a))
		}
		ch := s.board[candidate.x][candidate.y]
		s.a[k] = ch
		s.moves[k] = candidate
		nextNode := node.children[ch]
		s.backtrack(k, candidate, nextNode)
	}
}

func findWords(board [][]byte, words []string) []string {
	solution := NewSolution(board, words)
	root := solution.trie.root
	solution.backtrack(0, nil, root)
	return solution.result
}
