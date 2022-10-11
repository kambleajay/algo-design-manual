package list

type Node struct {
	val  int
	next *Node
}

type MyCircularQueue struct {
	head     *Node
	tail     *Node
	n        int
	capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{capacity: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &Node{val: value}
	if this.IsEmpty() {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.n++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = this.head.next
	this.n--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.val
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.n == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.n == this.capacity
}
