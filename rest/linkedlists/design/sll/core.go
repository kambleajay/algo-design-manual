package sll

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

func Constructor() MyLinkedList {
	head := &Node{}
	return MyLinkedList{head: head}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	x := this.head
	for i := 0; i < index+1; i++ {
		x = x.next
	}
	return x.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) nodeAtIndex(index int) *Node {
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.next
	}
	return pred
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	pred := this.nodeAtIndex(index)
	nodeToAdd := &Node{val: val, next: pred.next}
	pred.next = nodeToAdd
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	pred := this.nodeAtIndex(index)
	pred.next = pred.next.next
	this.size--
}
