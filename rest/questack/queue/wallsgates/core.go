package wallsgates

import "math"

type Point struct {
	row int
	col int
}

type Node struct {
	point *Point
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) enqueue(p *Point) {
	oldTail := q.tail
	q.tail = &Node{point: p}
	if q.isEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
	q.size++
}

func (q *Queue) dequeue() *Point {
	p := q.head.point
	q.head = q.head.next
	q.size--
	if q.isEmpty() {
		q.tail = nil
	}
	return p
}

const (
	EmptyRoom = math.MaxInt32
	Gate      = 0
)

func adj(p *Point, m, n int) []*Point {
	north := &Point{p.row - 1, p.col}
	east := &Point{p.row, p.col + 1}
	south := &Point{p.row + 1, p.col}
	west := &Point{p.row, p.col - 1}
	neighbours := []*Point{north, east, south, west}
	var result []*Point
	for _, nb := range neighbours {
		if nb.row >= 0 && nb.row < m && nb.col >= 0 && nb.col < n {
			result = append(result, nb)
		}
	}
	return result
}

func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 {
		return
	}
	m, n := len(rooms), len(rooms[0])
	q := Queue{}
	for i, row := range rooms {
		for j, item := range row {
			if item == Gate {
				q.enqueue(&Point{i, j})
			}
		}
	}
	for !q.isEmpty() {
		from := q.dequeue()
		for _, p := range adj(from, m, n) {
			if rooms[p.row][p.col] == EmptyRoom {
				rooms[p.row][p.col] = rooms[from.row][from.col] + 1
				q.enqueue(&Point{p.row, p.col})
			}
		}
	}
}
