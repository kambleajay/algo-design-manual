package bfs

type Node struct {
	sum  int
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

func (q *Queue) enqueue(n int) {
	oldTail := q.tail
	q.tail = &Node{sum: n}
	if q.isEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
	q.size++
}

func (q *Queue) dequeue() int {
	cb := q.head.sum
	q.head = q.head.next
	q.size--
	if q.isEmpty() {
		q.tail = nil
	}
	return cb
}

func (q *Queue) peek() int {
	if q.isEmpty() {
		return -1
	}
	return q.head.sum
}

func perfSquaresUpto(n int) []int {
	var result []int
	i := 1
	for {
		sq := i * i
		if sq > n {
			break
		} else {
			result = append(result, sq)
			i++
		}
	}
	return result
}

func numSquares(n int) int {
	squares := perfSquaresUpto(n)
	q := Queue{}
	seen := make(map[int]bool)
	for _, sq := range squares {
		q.enqueue(sq)
		seen[sq] = true
	}
	q.enqueue(-1)
	step := 1
	for !q.isEmpty() {
		num := q.dequeue()
		if num == -1 {
			step++
			if q.peek() != -1 {
				q.enqueue(-1)
			}
		} else if num == n {
			return step
		} else {
			for _, sq := range squares {
				newSum := num + sq
				if !seen[newSum] && newSum <= n {
					seen[newSum] = true
					q.enqueue(newSum)
				}
			}
		}
	}
	return -1
}
