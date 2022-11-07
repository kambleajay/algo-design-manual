package dfs

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

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

func (tst *TST) delete(key string) {
	tst.deleteRecur(tst.root, key, 0)
}

func (tst *TST) deleteRecur(x *Node, key string, d int) {
	ch := key[d]
	if x == nil {
		return
	}
	if ch < x.ch {
		tst.deleteRecur(x.left, key, d)
	} else if ch > x.ch {
		tst.deleteRecur(x.right, key, d)
	} else if d < len(key)-1 {
		tst.deleteRecur(x.mid, key, d+1)
	} else {
		x.val = false
	}
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

func between(n, min, max int) bool {
	return n >= min && n < max
}

func isValid(c Cell, m, n int) bool {
	return between(c.x, 0, m) && between(c.y, 0, n)
}

func adj(c Cell) []Cell {
	var nbs []Cell
	for _, dir := range dirs {
		x, y := c.x+dir[0], c.y+dir[1]
		nbs = append(nbs, Cell{x, y})
	}
	return nbs
}

func dfs(source Cell, s string, seen map[Cell]bool, trie TST, results map[string]bool, board [][]byte) {
	seen[source] = true
	m, n := len(board), len(board[0])
	if trie.contains(s) {
		results[s] = true
		trie.delete(s)
	}
	for _, nb := range adj(source) {
		if !seen[nb] && isValid(nb, m, n) {
			newStr := s + string(board[nb.x][nb.y])
			if trie.containsPrefix(newStr) {
				dfs(nb, newStr, seen, trie, results, board)
			}
		}
	}
	delete(seen, source)
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
	for i, rows := range board {
		for j, char := range rows {
			if firstChars[char] {
				cell := Cell{i, j}
				dfs(cell, string(char), make(map[Cell]bool), trie, resultSet, board)
			}
		}
	}
	for k, _ := range resultSet {
		result = append(result, k)
	}
	return result
}
