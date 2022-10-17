package stackwithques

type Node struct {
	val  int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}

func (q *Queue) enqueue(val int) {
	oldTail := q.tail
	q.tail = &Node{val: val}
	if q.isEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
	q.size++
}

func (q *Queue) dequeue() int {
	item := q.head.val
	q.head = q.head.next
	if q.isEmpty() {
		q.tail = nil
	}
	q.size--
	return item
}

func (q *Queue) top() int {
	return q.head.val
}

type MyStack struct {
	q1  *Queue
	q2  *Queue
	top int
}

func Constructor() MyStack {
	q1, q2 := &Queue{}, &Queue{}
	return MyStack{q1: q1, q2: q2}
}

func (this *MyStack) Push(x int) {
	this.q2.enqueue(x)
	this.top = x
	for !this.q1.isEmpty() {
		this.q2.enqueue(this.q1.dequeue())
	}
	this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack) Pop() int {
	item := this.q1.dequeue()
	if !this.q1.isEmpty() {
		this.top = this.q1.top()
	}
	return item
}

func (this *MyStack) Top() int {
	return this.top
}

func (this *MyStack) Empty() bool {
	return this.q1.isEmpty()
}
