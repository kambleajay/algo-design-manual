package list

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
	size int
}

func NewList() *List {
	head := &Node{}
	return &List{head: head}
}

func (list *List) Add(n int) {
	pred, succ := list.head, list.head.next
	node := &Node{val: n}
	pred.next = node
	node.next = succ
	list.size++
}

func (list *List) Remove(n int) {
	pred, succ := list.head, list.head.next
	for succ != nil {
		if succ.val == n {
			break
		}
		pred = succ
		succ = succ.next
	}
	pred.next = succ.next
	list.size--
}

func (list *List) Contains(n int) bool {
	if list.size == 0 {
		return false
	}
	x := list.head.next
	for x != nil {
		if x.val == n {
			return true
		}
	}
	return false
}

type MyHashSet struct {
	items []*List
	q     int
}

func Constructor() MyHashSet {
	q := 997
	items := make([]*List, q)
	for i := 0; i < len(items); i++ {
		items[i] = NewList()
	}
	return MyHashSet{items, q}
}

func (this *MyHashSet) hash(key int) int {
	return key % this.q
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	index := this.hash(key)
	this.items[index].Add(key)
}

func (this *MyHashSet) Remove(key int) {
	if !this.Contains(key) {
		return
	}
	index := this.hash(key)
	this.items[index].Remove(key)
}

func (this *MyHashSet) Contains(key int) bool {
	index := this.hash(key)
	return this.items[index].Contains(key)
}
