package array

type MyCircularQueue struct {
	head     int
	items    []int
	n        int
	capacity int
}

func Constructor(k int) MyCircularQueue {
	a := make([]int, k)
	return MyCircularQueue{items: a, capacity: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	tail := (this.head + this.n) % this.capacity
	this.items[tail] = value
	this.n++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.capacity
	this.n--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	tail := (this.head + this.n - 1) % this.capacity
	return this.items[tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.n == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.n == this.capacity
}
