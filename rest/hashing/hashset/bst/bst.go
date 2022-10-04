package bst

const Prime = 997

const (
	RED   = true
	BLACK = false
)

type Node struct {
	key   int
	left  *Node
	right *Node
	color bool
}

type RedBlackBST struct {
	root *Node
}

func NewRedBlackBST() *RedBlackBST {
	return &RedBlackBST{}
}

func (tree *RedBlackBST) isEmpty() bool {
	return tree.root == nil
}

func isRed(h *Node) bool {
	if h == nil {
		return false
	}
	return h.color == RED
}

func isBlack(h *Node) bool {
	return !isRed(h)
}

func rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	return x
}

func rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	return x
}

func flipColors(h *Node) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

func balance(h *Node) *Node {
	if isRed(h.right) && isBlack(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}
	return h
}

func put(h *Node, key int) *Node {
	if h == nil {
		return &Node{key: key, color: RED}
	}
	if key < h.key {
		h.left = put(h.left, key)
	} else if key > h.key {
		h.right = put(h.right, key)
	} else {
		h.key = key
	}
	return balance(h)
}

func (tree *RedBlackBST) put(key int) {
	tree.root = put(tree.root, key)
	tree.root.color = BLACK
}

func (tree *RedBlackBST) get(key int) *Node {
	x := tree.root
	for x != nil {
		if key < x.key {
			x = x.left
		} else if key > x.key {
			x = x.right
		} else {
			return x
		}
	}
	return nil
}

func (tree *RedBlackBST) contains(key int) bool {
	return tree.get(key) != nil
}

func moveRedLeft(h *Node) *Node {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}
	return h
}

func moveRedRight(h *Node) *Node {
	flipColors(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}
	return h
}

func min(h *Node) *Node {
	if h.left == nil {
		return h
	}
	return min(h.left)
}

func deleteMin(h *Node) *Node {
	if h.left == nil {
		return nil
	}
	if isBlack(h.left) && isBlack(h.left.left) {
		h = moveRedLeft(h)
	}
	h.left = deleteMin(h.left)
	return balance(h)
}

func delete(h *Node, key int) *Node {
	if key < h.key {
		if isBlack(h.left) && isBlack(h.left.left) {
			h = moveRedLeft(h)
		}
		h.left = delete(h.left, key)
	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}
		if key == h.key && h.right == nil {
			return nil
		}
		if isBlack(h.right) && isBlack(h.right.left) {
			h = moveRedRight(h)
		}
		if key == h.key {
			x := min(h.right)
			h.key = x.key
			h.right = deleteMin(h.right)
		} else {
			h.right = delete(h.right, key)
		}
	}
	return balance(h)
}

func (tree *RedBlackBST) delete(key int) {
	if !tree.contains(key) {
		return
	}
	if isBlack(tree.root.left) && isBlack(tree.root.right) {
		tree.root.color = RED
	}
	tree.root = delete(tree.root, key)
	if !tree.isEmpty() {
		tree.root.color = BLACK
	}
}

type MyHashSet struct {
	items [Prime]*RedBlackBST
}

func Constructor() MyHashSet {
	set := MyHashSet{}
	for i := 0; i < Prime; i++ {
		set.items[i] = NewRedBlackBST()
	}
	return set
}

func (this *MyHashSet) hash(key int) int {
	return key % Prime
}

func (this *MyHashSet) Add(key int) {
	index := this.hash(key)
	this.items[index].put(key)
}

func (this *MyHashSet) Remove(key int) {
	index := this.hash(key)
	this.items[index].delete(key)
}

func (this *MyHashSet) Contains(key int) bool {
	index := this.hash(key)
	return this.items[index].contains(key)
}
