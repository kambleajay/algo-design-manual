package serialize

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	if tn == nil {
		return ""
	}
	return fmt.Sprintf("%s{%d}%s", tn.Left, tn.Val, tn.Right)
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func makeString(node *TreeNode, k int, m []string) {
	if k > len(m)-1 {
		return
	}
	if node == nil {
		m[k] = "#"
		return
	}
	m[k] = strconv.Itoa(node.Val)
	makeString(node.Left, 2*k, m)
	makeString(node.Right, (2*k)+1, m)
}

type NodeAndDepth struct {
	node *TreeNode
	d    int
}

type Queue struct {
	items []NodeAndDepth
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(node NodeAndDepth) {
	q.items = append(q.items, node)
}

func (q *Queue) dequeue() NodeAndDepth {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q Queue
	var nd NodeAndDepth
	q.enqueue(NodeAndDepth{root, 1})
	for !q.isEmpty() {
		nd = q.dequeue()
		node := nd.node
		if node.Left != nil {
			q.enqueue(NodeAndDepth{node.Left, nd.d + 1})
		}
		if node.Right != nil {
			q.enqueue(NodeAndDepth{node.Right, nd.d + 1})
		}
	}
	return nd.d
}

func size(d int) int {
	return int(math.Pow(2, float64(d))) - 1
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	d := depth(root)
	n := size(d)
	indexToNode := make([]string, n+1)
	for i := 0; i < len(indexToNode); i++ {
		indexToNode[i] = "#"
	}
	makeString(root, 1, indexToNode)
	var result []string
	for i := 1; i < len(indexToNode); i++ {
		result = append(result, indexToNode[i])
	}
	return strings.Join(result, ",")
}

func makeTree(ss []string, k int) *TreeNode {
	if k > len(ss)-1 {
		return nil
	}
	ch := ss[k]
	if ch == "#" {
		return nil
	}
	val, _ := strconv.Atoi(ch)
	node := &TreeNode{Val: val}
	node.Left = makeTree(ss, 2*k)
	node.Right = makeTree(ss, (2*k)+1)
	return node
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	ss := []string{"0"}
	ss = append(ss, strings.Split(data, ",")...)
	return makeTree(ss, 1)
}
