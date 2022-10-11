package circular

import (
	"algo/rest/questack/queue/circular/array"
	"algo/rest/questack/queue/circular/list"
	"testing"
)

func arrayQueue(k int) array.MyCircularQueue {
	return array.Constructor(k)
}

func listQueue(k int) list.MyCircularQueue {
	return list.Constructor(k)
}

type Queue struct {
	size int
}

type Enqueue struct {
	n    int
	want bool
}

type Rear struct {
	want int
}

type IsFull struct {
	want bool
}

type Dequeue struct {
	want bool
}

type Front struct {
	want int
}

type IsEmpty struct {
	want bool
}

func TestCircularQueue(t *testing.T) {
	var tests = [][]interface{}{
		{Queue{3}, Enqueue{1, true}, Enqueue{2, true}, Enqueue{3, true},
			Enqueue{4, false}, Rear{3}, IsFull{true}, Dequeue{true},
			Enqueue{4, true}, Rear{4}, IsEmpty{false}, Front{2}, IsFull{true},
			Dequeue{true}, Dequeue{true}, Dequeue{true}, Dequeue{false}, IsEmpty{true},
			Front{-1}, Rear{-1}},
		{Queue{3}, Enqueue{2, true}, Rear{2}, Front{2}},
		{Queue{3}, Enqueue{1, true}, Enqueue{2, true}, Enqueue{3, true},
			Enqueue{4, false}, Front{1}, Rear{3}, IsFull{true}, IsEmpty{false},
			Dequeue{true}, IsFull{false}, IsEmpty{false}, Front{2}, Rear{3},
			Dequeue{true}, IsFull{false}, IsEmpty{false}, Front{3}, Rear{3},
			Dequeue{true}, IsEmpty{true}, IsFull{false}, Front{-1}, Rear{-1}},
		{Queue{81}, Enqueue{69, true}, Dequeue{true}, Enqueue{92, true},
			Enqueue{12, true}, Dequeue{true}, IsFull{false}, IsFull{false}, Front{12}},
	}
	for i, test := range tests {
		var q list.MyCircularQueue
		for _, command := range test {
			switch command := command.(type) {
			case Queue:
				q = listQueue(command.size)
			case Enqueue:
				if got := q.EnQueue(command.n); got != command.want {
					t.Errorf("[%d] Enqueue(%d) = %t, want %t", i+1, command.n, got, command.want)
				}
			case Rear:
				if got := q.Rear(); got != command.want {
					t.Errorf("[%d] Rear() = %d, want %d", i+1, got, command.want)
				}
			case IsFull:
				if got := q.IsFull(); got != command.want {
					t.Errorf("[%d] IsFull() = %t, want %t", i+1, got, command.want)
				}
			case Dequeue:
				if got := q.DeQueue(); got != command.want {
					t.Errorf("[%d] Dequeue() = %t, want %t", i+1, got, command.want)
				}
			case IsEmpty:
				if got := q.IsEmpty(); got != command.want {
					t.Errorf("[%d] IsEmpty() = %t, want %t", i+1, got, command.want)
				}
			case Front:
				if got := q.Front(); got != command.want {
					t.Errorf("[%d] Front() = %d, want %d", i+1, got, command.want)
				}
			}
		}
	}
}
