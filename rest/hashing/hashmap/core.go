package hashmap

const Prime = 997

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Node struct {
	key    int
	val    int
	height int
	left   *Node
	right  *Node
}

type AVLTree struct {
	root *Node
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func height(x *Node) int {
	if x == nil {
		return -1
	}
	return x.height
}

func calculateHeight(x *Node) int {
	return 1 + max(height(x.left), height(x.right))
}

func balanceFactor(x *Node) int {
	return height(x.right) - height(x.left)
}

func rotateLeft(x *Node) *Node {
	y := x.right
	x.right = y.left
	y.left = x
	x.height = calculateHeight(x)
	y.height = calculateHeight(y)
	return y
}

func rotateRight(x *Node) *Node {
	y := x.left
	x.left = y.right
	y.right = x
	x.height = calculateHeight(x)
	y.height = calculateHeight(y)
	return y
}

func balance(x *Node) *Node {
	if balanceFactor(x) > 1 {
		if balanceFactor(x.right) < 0 {
			x.right = rotateRight(x.right)
		}
		x = rotateLeft(x)
	} else if balanceFactor(x) < -1 {
		if balanceFactor(x.left) > 0 {
			x.left = rotateLeft(x.left)
		}
		x = rotateRight(x)
	}
	return x
}

func put(x *Node, key, val int) *Node {
	if x == nil {
		return &Node{key: key, val: val}
	}
	if key < x.key {
		x.left = put(x.left, key, val)
	} else if key > x.key {
		x.right = put(x.right, key, val)
	} else {
		x.val = val
		return x
	}
	x.height = calculateHeight(x)
	return balance(x)
}

func (tree *AVLTree) put(key int, val int) {
	tree.root = put(tree.root, key, val)
}

func get(x *Node, key int) *Node {
	if x == nil {
		return nil
	}
	if key < x.key {
		return get(x.left, key)
	} else if key > x.key {
		return get(x.right, key)
	} else {
		return x
	}
}

func (tree *AVLTree) get(key int) int {
	r := get(tree.root, key)
	if r == nil {
		return -1
	}
	return r.val
}

func (tree *AVLTree) contains(key int) bool {
	return tree.get(key) != -1
}

func min(x *Node) *Node {
	if x.left == nil {
		return x
	}
	return min(x.left)
}

func deleteMin(x *Node) *Node {
	if x.left == nil {
		return x.right
	}
	x.left = deleteMin(x.left)
	x.height = calculateHeight(x)
	return balance(x)
}

func delete(x *Node, key int) *Node {
	if key < x.key {
		x.left = delete(x.left, key)
	} else if key > x.key {
		x.right = delete(x.right, key)
	} else {
		if x.left == nil {
			return x.right
		} else if x.right == nil {
			return x.left
		} else {
			y := x
			x = min(y.right)
			x.right = deleteMin(y.right)
			x.left = y.left
		}
	}
	x.height = calculateHeight(x)
	return balance(x)
}

func (tree *AVLTree) delete(key int) {
	if !tree.contains(key) {
		return
	}
	tree.root = delete(tree.root, key)
}

type MyHashMap struct {
	items [Prime]*AVLTree
}

func Constructor() MyHashMap {
	m := MyHashMap{}
	for i := 0; i < Prime; i++ {
		m.items[i] = NewAVLTree()
	}
	return m
}

func hash(key int) int {
	return key % Prime
}

func (this *MyHashMap) Put(key int, value int) {
	index := hash(key)
	this.items[index].put(key, value)
}

func (this *MyHashMap) Get(key int) int {
	index := hash(key)
	return this.items[index].get(key)
}

func (this *MyHashMap) Remove(key int) {
	index := hash(key)
	this.items[index].delete(key)
}
