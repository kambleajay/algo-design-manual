package temp

func findDayOfWarmerTemp(temp int, tempDay int, tempToDay map[int][]int) int {
	var dayOfWarmerTemp int
	for j := temp + 1; j <= 100; j++ {
		days, ok := tempToDay[j]
		if ok {
			for _, day := range days {
				diffOfDays := day - tempDay
				if day > tempDay && (dayOfWarmerTemp == 0 || diffOfDays < dayOfWarmerTemp) {
					dayOfWarmerTemp = diffOfDays
					if dayOfWarmerTemp == 1 {
						return dayOfWarmerTemp
					}
				}
			}
		}
	}
	return dayOfWarmerTemp
}

func dailyTemperatures(temperatures []int) []int {
	tempToDay := make(map[int][]int)
	for i, temp := range temperatures {
		tempToDay[temp] = append(tempToDay[temp], i)
	}
	answer := make([]int, len(temperatures))
	for i, temp := range temperatures {
		answer[i] = findDayOfWarmerTemp(temp, i, tempToDay)
	}
	return answer
}

const (
	RED   = true
	BLACK = false
)

type Node struct {
	key   int
	val   int
	left  *Node
	right *Node
	color bool
}

type RedBlackBST struct {
	root *Node
}

func (_ *RedBlackBST) isRed(x *Node) bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

func (_ *RedBlackBST) rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	return x
}

func (_ *RedBlackBST) rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
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
		if x.key < k {
			x = x.left
		} else if x.key > k {
			x = x.right
		} else {
			return x.val, true
		}
	}
	return 0, false
}

func (t *RedBlackBST) Put(k int, v int) {
	t.root = t.put(t.root, k, v)
	t.root.color = BLACK
}

func (t *RedBlackBST) put(x *Node, k int, v int) *Node {
	if x == nil {
		return &Node{key: k, val: v, color: RED}
	}
	if x.key < k {
		x.left = t.put(x.left, k, v)
	} else if x.key > k {
		x.right = t.put(x.right, k, v)
	} else {
		x.val = v
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

	return x
}
