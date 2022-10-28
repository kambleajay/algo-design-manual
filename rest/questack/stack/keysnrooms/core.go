package keysnrooms

type Node struct {
	val  int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
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
}

func (q *Queue) dequeue() int {
	item := q.head.val
	q.head = q.head.next
	if q.isEmpty() {
		q.tail = nil
	}
	return item
}

func bfs(rooms [][]int, source int, seen map[int]bool) {
	queue := Queue{}
	queue.enqueue(source)
	seen[source] = true
	for !queue.isEmpty() {
		v := queue.dequeue()
		for _, w := range rooms[v] {
			if !seen[w] {
				queue.enqueue(w)
				seen[w] = true
			}
		}
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	seen := make(map[int]bool)
	bfs(rooms, 0, seen)
	return len(seen) == len(rooms)
}
