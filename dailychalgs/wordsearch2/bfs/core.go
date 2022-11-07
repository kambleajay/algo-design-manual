package dfs

type Node struct {
	val              bool
	ch               byte
	left, mid, right *Node
}

type TST struct {
	root *Node
}

func (tst *TST) put(key string) {
	tst.root = tst.putRecur(tst.root, key, 0)
}

func (tst *TST) putRecur(x *Node, key string, d int) *Node {
	ch := key[d]
	if x == nil {
		x = &Node{ch: ch}
	}
	if ch < x.ch {
		x.left = tst.putRecur(x.left, key, d)
	} else if ch > x.ch {
		x.right = tst.putRecur(x.right, key, d)
	} else if d < len(key)-1 {
		x.mid = tst.putRecur(x.mid, key, d+1)
	} else {
		x.val = true
	}
	return x
}

func (tst *TST) get(key string) bool {
	x := tst.getRecur(tst.root, key, 0)
	return x != nil && x.val
}

func (tst *TST) getRecur(x *Node, key string, d int) *Node {
	ch := key[d]
	if x == nil {
		return nil
	}
	if ch < x.ch {
		return tst.getRecur(x.left, key, d)
	} else if ch > x.ch {
		return tst.getRecur(x.right, key, d)
	} else if d < len(key)-1 {
		return tst.getRecur(x.mid, key, d+1)
	} else {
		return x
	}
}

func (tst *TST) contains(key string) bool {
	return tst.get(key)
}

func (tst *TST) containsPrefix(key string) bool {
	x := tst.getRecur(tst.root, key, 0)
	return x != nil
}

type Cell struct {
	x, y int
}

type State struct {
	Cell
	s     string
	paths map[Cell]bool
}

type Queue struct {
	items []State
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(s State) {
	q.items = append(q.items, s)
}

func (q *Queue) dequeue() State {
	s := q.items[0]
	q.items = q.items[1:]
	return s
}

func between(n, min, max int) bool {
	return n >= min && n < max
}

func next(from State, board [][]byte) []State {
	var states []State
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	m, n := len(board), len(board[0])
	for _, dir := range dirs {
		x, y := from.x+dir[0], from.y+dir[1]
		c := Cell{x: x, y: y}
		if between(x, 0, m) && between(y, 0, n) && !from.paths[c] {
			s := from.s + string(board[x][y])
			m := make(map[Cell]bool)
			for k, v := range from.paths {
				m[k] = v
			}
			m[c] = true
			states = append(states, State{c, s, m})
		}
	}
	return states
}

func findWords(board [][]byte, words []string) []string {
	var result []string
	resultSet := make(map[string]bool)
	trie := TST{}
	firstChars := make(map[byte]bool)
	for _, word := range words {
		trie.put(word)
		firstChars[word[0]] = true
	}
	q := Queue{}
	for i, rows := range board {
		for j, char := range rows {
			if firstChars[char] {
				cell := Cell{i, j}
				m := map[Cell]bool{cell: true}
				state := State{cell, string(char), m}
				q.enqueue(state)
			}
		}
	}
	for !q.isEmpty() {
		from := q.dequeue()
		if trie.contains(from.s) {
			resultSet[from.s] = true
		}
		for _, to := range next(from, board) {
			if trie.containsPrefix(to.s) {
				q.enqueue(to)
			}
		}
	}
	for k, _ := range resultSet {
		result = append(result, k)
	}
	return result
}
