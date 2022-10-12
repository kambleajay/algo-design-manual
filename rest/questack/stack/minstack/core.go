package minstack

type MinStack struct {
	stack [][]int
}

func Constructor() MinStack {
	return MinStack{}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (this *MinStack) isEmpty() bool {
	return len(this.stack) == 0
}

func (this *MinStack) Push(val int) {
	if this.isEmpty() {
		this.stack = append(this.stack, []int{val, min(val, val)})
	} else {
		currentMin := this.GetMin()
		this.stack = append(this.stack, []int{val, min(currentMin, val)})
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1][1]
}
