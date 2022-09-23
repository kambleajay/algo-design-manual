package dll

type Node struct {
	val  int
	next *Node
	prev *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() MyLinkedList {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return MyLinkedList{head: head, tail: tail}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	var x *Node
	if index+1 < this.size-index {
		x = this.head
		for i := 0; i < index+1; i++ {
			x = x.next
		}
	} else {
		x = this.tail
		for i := 0; i < this.size-index; i++ {
			x = x.prev
		}
	}
	return x.val
}

func (this *MyLinkedList) addBetween(val int, pred *Node, succ *Node) {
	nodeToAdd := &Node{val: val}
	pred.next = nodeToAdd
	nodeToAdd.next = succ
	succ.prev = nodeToAdd
	nodeToAdd.prev = pred
	this.size++
}

func (this *MyLinkedList) AddAtHead(val int) {
	pred, succ := this.head, this.head.next
	this.addBetween(val, pred, succ)
}

func (this *MyLinkedList) AddAtTail(val int) {
	pred, succ := this.tail.prev, this.tail
	this.addBetween(val, pred, succ)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	var pred, succ *Node
	if index < this.size-index {
		pred = this.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next
	} else {
		succ = this.tail
		for i := 0; i < this.size-index; i++ {
			succ = succ.prev
		}
		pred = succ.prev
	}
	this.addBetween(val, pred, succ)
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	var pred, succ *Node
	if index < this.size-index {
		pred = this.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next.next
	} else {
		succ = this.tail
		for i := 0; i < this.size-index-1; i++ {
			succ = succ.prev
		}
		pred = succ.prev.prev
	}
	pred.next = succ
	succ.prev = pred
	this.size--
}
