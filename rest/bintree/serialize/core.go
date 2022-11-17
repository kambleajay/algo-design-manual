package serialize

import (
	"fmt"
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

func serializeHelper(node *TreeNode, sa *[]string) {
	if node == nil {
		*sa = append(*sa, "#")
		return
	}
	*sa = append(*sa, strconv.Itoa(node.Val))
	serializeHelper(node.Left, sa)
	serializeHelper(node.Right, sa)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	var sa []string
	serializeHelper(root, &sa)
	return strings.Join(sa, ",")
}

type Queue struct {
	items []string
}

func NewQueue(items []string) *Queue {
	return &Queue{items: items}
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(s string) {
	q.items = append(q.items, s)
}

func (q *Queue) dequeue() string {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) top() string {
	return q.items[0]
}

func deserializeHelper(q *Queue) *TreeNode {
	if q.top() == "#" {
		q.dequeue()
		return nil
	}
	val, _ := strconv.Atoi(q.top())
	node := &TreeNode{Val: val}
	q.dequeue()
	node.Left = deserializeHelper(q)
	node.Right = deserializeHelper(q)
	return node
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sa := strings.Split(data, ",")
	q := NewQueue(sa)
	return deserializeHelper(q)
}
