package temp

import (
	"fmt"
)

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

func dailyTemperatures1(temperatures []int) []int {
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
	vals  []int
	left  *Node
	right *Node
	color bool
}

func (x *Node) String() string {
	return fmt.Sprintf("[%d/%v]", x.key, x.vals)
}

func compareTo(k int, x *Node) int {
	if k < x.key {
		return -1
	} else if k > x.key {
		return 1
	}
	return 0
}

type RedBlackBST struct {
	root *Node
}

func StringRecur(x *Node) string {
	if x == nil {
		return ""
	}
	return fmt.Sprintf("%s-[%s]-%s", StringRecur(x.left), x, StringRecur(x.right))
}

func (t *RedBlackBST) String() string {
	return StringRecur(t.root)
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

func (t *RedBlackBST) Get(k int) ([]int, bool) {
	x := t.root
	for x != nil {
		cmp := compareTo(k, x)
		if cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x.vals, true
		}
	}
	return nil, false
}

func (t *RedBlackBST) Put(k int, v int) {
	t.root = t.put(t.root, k, v)
	t.root.color = BLACK
}

func (t *RedBlackBST) put(x *Node, k int, v int) *Node {
	if x == nil {
		return &Node{key: k, vals: []int{v}, color: RED}
	}
	cmp := compareTo(k, x)
	if cmp < 0 {
		x.left = t.put(x.left, k, v)
	} else if cmp > 0 {
		x.right = t.put(x.right, k, v)
	} else {
		x.vals = append(x.vals, v)
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

func (t *RedBlackBST) nodes(lo int, hi int) []*Node {
	var nodes []*Node
	t.nodesRecur(lo, hi, t.root, &nodes)
	return nodes
}

func (t *RedBlackBST) nodesRecur(lo int, hi int, x *Node, nodes *[]*Node) {
	if x == nil {
		return
	}
	cmpLo := compareTo(lo, x)
	cmpHi := compareTo(hi, x)
	if cmpLo < 0 {
		t.nodesRecur(lo, hi, x.left, nodes)
	}
	if cmpLo <= 0 && cmpHi >= 0 {
		*nodes = append(*nodes, x)
	}
	if cmpHi > 0 {
		t.nodesRecur(lo, hi, x.right, nodes)
	}
}

func (t *RedBlackBST) FindDayOfWarmerTemp(temp int, day int) int {
	nodes := t.nodes(temp+1, 100)
	earliestDay := 0
	for _, node := range nodes {
		days := node.vals
		for _, d := range days {
			diffInDays := d - day
			if earliestDay == 0 {
				if diffInDays > 0 {
					earliestDay = diffInDays
					if earliestDay == 1 {
						return earliestDay
					}
				}
			} else {
				if diffInDays > 0 && diffInDays < earliestDay {
					earliestDay = diffInDays
					if earliestDay == 1 {
						return earliestDay
					}
				}
			}
		}
	}
	return earliestDay
}

func dailyTemperatures2(temperatures []int) []int {
	tempTree := RedBlackBST{}
	for day, temp := range temperatures {
		tempTree.Put(temp, day)
	}
	answer := make([]int, len(temperatures))
	for day, temp := range temperatures {
		answer[day] = tempTree.FindDayOfWarmerTemp(temp, day)
	}
	return answer
}
