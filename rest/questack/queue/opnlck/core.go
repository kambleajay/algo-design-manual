package opnlck

import "fmt"

type Combination struct {
	wheels [4]int
}

func (cb *Combination) String() string {
	return fmt.Sprintf("%d%d%d%d", cb.wheels[0], cb.wheels[1], cb.wheels[2], cb.wheels[3])
}

type Node struct {
	cb   *Combination
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

func (q *Queue) enqueue(cb *Combination) {
	oldTail := q.tail
	q.tail = &Node{cb: cb}
	if q.isEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
	q.size++
}

func (q *Queue) dequeue() *Combination {
	cb := q.head.cb
	q.head = q.head.next
	q.size--
	if q.isEmpty() {
		q.tail = nil
	}
	return cb
}

func (q *Queue) peek() *Combination {
	if q.isEmpty() {
		return nil
	}
	return q.head.cb
}

func asCombination(s string) Combination {
	var wheels [4]int
	for i, r := range s {
		wheels[i] = int(r - '0')
	}
	return Combination{wheels}
}

func adj(cb *Combination) []*Combination {
	var cbs []*Combination
	options := map[int][]int{0: {1, 9}, 1: {2, 0}, 2: {3, 1}, 3: {4, 2}, 4: {5, 3}, 5: {6, 4}, 6: {7, 5}, 7: {8, 6}, 8: {9, 7}, 9: {0, 8}}
	for i := 0; i < 4; i++ {
		nextPrev := options[cb.wheels[i]]
		cbNext := &Combination{wheels: cb.wheels}
		cbNext.wheels[i] = nextPrev[0]
		cbPrev := &Combination{wheels: cb.wheels}
		cbPrev.wheels[i] = nextPrev[1]
		cbs = append(cbs, cbNext, cbPrev)
	}
	return cbs
}

func openLock(deadends []string, target string) int {
	sourceStr := "0000"
	source := asCombination(sourceStr)
	seen := make(map[Combination]bool)
	q := Queue{}
	q.enqueue(&source)
	q.enqueue(nil)
	seen[source] = true
	var step int
	dead := make(map[string]bool)
	for _, de := range deadends {
		dead[de] = true
	}
	for !q.isEmpty() {
		cb := q.dequeue()
		if cb == nil {
			step++
			if q.peek() != nil {
				q.enqueue(nil)
			}
		} else if cb.String() == target {
			return step
		} else if !dead[cb.String()] {
			for _, nb := range adj(cb) {
				if !seen[*nb] {
					q.enqueue(nb)
					seen[*nb] = true
				}
			}
		}
	}
	return -1
}
