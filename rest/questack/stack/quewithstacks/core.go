package quewithstacks

type Stack struct {
	items []int
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) peek() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) pop() int {
	item := s.peek()
	s.items = s.items[:len(s.items)-1]
	return item
}

type MyQueue struct {
	stack   *Stack
	asQueue *Stack
}

func Constructor() MyQueue {
	return MyQueue{&Stack{}, &Stack{}}
}

func (this *MyQueue) Push(x int) {
	this.stack.push(x)
}

func (this *MyQueue) copyStackToQueue() {
	for !this.stack.isEmpty() {
		item := this.stack.pop()
		this.asQueue.push(item)
	}
}

func (this *MyQueue) balance() {
	if this.asQueue.isEmpty() {
		this.copyStackToQueue()
	}
}

func (this *MyQueue) Peek() int {
	this.balance()
	return this.asQueue.peek()
}

func (this *MyQueue) Pop() int {
	this.balance()
	return this.asQueue.pop()
}

func (this *MyQueue) Empty() bool {
	return this.stack.isEmpty() && this.asQueue.isEmpty()
}
