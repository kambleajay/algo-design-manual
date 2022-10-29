package zeroonemat

type Point struct {
	x int
	y int
}

type Node struct {
	p    *Point
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}

func (q *Queue) enqueue(p *Point) {
	oldTail := q.tail
	q.tail = &Node{p: p}
	if q.isEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
}

func (q *Queue) dequeue() *Point {
	item := q.head.p
	q.head = q.head.next
	if q.isEmpty() {
		q.tail = nil
	}
	return item
}

func between(n, min, max int) bool {
	return n >= min && n < max
}

func adj(p *Point, m, n int) []*Point {
	var nbs []*Point
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, dir := range dirs {
		nb := &Point{p.x + dir[0], p.y + dir[1]}
		if between(nb.x, 0, m) && between(nb.y, 0, n) {
			nbs = append(nbs, nb)
		}
	}
	return nbs
}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	seen := make(map[Point]bool)
	result := make([][]int, m)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, n)
	}
	q := Queue{}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				source := Point{i, j}
				q.enqueue(&source)
				seen[source] = true
			}
		}
	}
	for !q.isEmpty() {
		v := q.dequeue()
		for _, w := range adj(v, m, n) {
			if !seen[*w] && mat[w.x][w.y] == 1 {
				result[w.x][w.y] = result[v.x][v.y] + 1
				q.enqueue(w)
				seen[*w] = true
			}
		}
	}
	return result
}
