package movavg

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
	return q.size == 0
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
	q.size--
	if q.isEmpty() {
		q.tail = nil
	}
	return item
}

type MovingAverage struct {
	queue *Queue
	sum   int
	size  int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{queue: &Queue{}, size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	this.queue.enqueue(val)
	this.sum += val
	if this.queue.size > this.size {
		item := this.queue.dequeue()
		this.sum -= item
	}
	return float64(this.sum) / float64(this.queue.size)
}
