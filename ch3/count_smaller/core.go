package count_smaller

const (
	RED   = true
	BLACK = false
)

type Node struct {
	key, val, size, freq int
	left                 *Node
	right                *Node
	color                bool
}

type RedBlackBST struct {
	root *Node
}

func (t *RedBlackBST) Size() int {
	return t.size(t.root)
}

func (_ *RedBlackBST) size(x *Node) int {
	if x == nil {
		return 0
	}
	return x.size
}

func (_ *RedBlackBST) sizeOfNode(x *Node) int {
	if x.freq > 1 {
		return x.freq
	}
	return 1
}

func (t *RedBlackBST) Rank(k int) int {
	return t.rank(k, t.root)
}

func (t *RedBlackBST) rank(k int, x *Node) int {
	if x == nil {
		return 0
	}
	if k < x.key {
		return t.rank(k, x.left)
	} else if k > x.key {
		return t.sizeOfNode(x) + t.size(x.left) + t.rank(k, x.right)
	} else {
		return t.size(x.left)
	}
}

func (_ *RedBlackBST) isRed(x *Node) bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

func (t *RedBlackBST) rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	x.size = h.size
	h.size = t.sizeOfNode(h) + t.size(h.left) + t.size(h.right)
	return x
}

func (t *RedBlackBST) rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	x.size = h.size
	h.size = t.sizeOfNode(h) + t.size(h.left) + t.size(h.right)
	return x
}

func (_ *RedBlackBST) flipColors(h *Node) {
	h.color = RED
	h.left.color = BLACK
	h.right.color = BLACK
}

func (t *RedBlackBST) Get(k int) (int, bool) {
	x := t.root
	for x != nil {
		if k < x.key {
			x = x.left
		} else if k > x.key {
			x = x.right
		} else {
			return x.val, true
		}
	}
	return -1, false
}

func (t *RedBlackBST) Put(k, v int) {
	t.root = t.put(t.root, k, v)
	t.root.color = BLACK
}

func (t *RedBlackBST) put(x *Node, k, v int) *Node {
	if x == nil {
		return &Node{key: k, val: v, color: RED, size: 1, freq: 1}
	}
	if k < x.key {
		x.left = t.put(x.left, k, v)
	} else if k > x.key {
		x.right = t.put(x.right, k, v)
	} else {
		x.val = v
		x.freq = x.freq + 1
	}

	if t.isRed(x.right) && !t.isRed(x.left) {
		x = t.rotateLeft(x)
	}
	if t.isRed(x.left) && t.isRed(x.left.left) {
		x = t.rotateRight(x)
	}
	if t.isRed(x.left) && t.isRed(x.right) {
		t.flipColors(x)
	}

	x.size = t.sizeOfNode(x) + t.size(x.left) + t.size(x.right)
	return x
}

func countSmaller1(nums []int) []int {
	var answer []int
	var count int
	for i := 0; i < len(nums); i++ {
		count = 0
		for j := i + 1; j < len(nums) && j > i; j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		answer = append(answer, count)
	}
	return answer
}

func countSmaller2(nums []int) []int {
	answer := make([]int, len(nums))
	bst := RedBlackBST{}
	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			answer[j] = 0
		} else {
			answer[j] = bst.Rank(nums[j])
		}
		bst.Put(nums[j], nums[j])
	}
	return answer
}

func query(left, right int, tree []int, size int) int {
	l, r := left+size, right+size
	var count int
	for l < r {
		if l%2 == 1 {
			count += tree[l]
			l++
		}
		if r%2 == 1 {
			r--
			count += tree[r]
		}
		l /= 2
		r /= 2
	}
	return count
}

func update(index, val int, tree []int, size int) {
	k := index + size
	tree[k] += val
	for k > 1 {
		k /= 2
		c1, c2 := tree[k*2], tree[(k*2)+1]
		tree[k] = c1 + c2
	}
}

func countSmaller3(nums []int) []int {
	offset, size := 10000, (10000*2)+1
	tree := make([]int, size*2)
	answer := make([]int, len(nums))
	for j := len(nums) - 1; j >= 0; j-- {
		count := query(0, nums[j]+offset, tree, size)
		answer[j] = count
		update(nums[j]+offset, 1, tree, size)
	}
	return answer
}

func countSmaller(nums []int) []int {
	return countSmaller3(nums)
}
