package astar

type Cell struct {
	x int
	y int
}

type Step struct {
	Cell
	obstCount int
}

type State struct {
	Step
	dist     int
	estimate int
}

func NewState(x, y, obstCount, dist int, target Cell) *State {
	manhattanDist := target.x - x + target.y - y
	est := manhattanDist + dist
	return &State{Step{Cell{x, y}, obstCount}, dist, est}
}

type MinPQ struct {
	items []*State
	n     int
}

func NewMinPQ() MinPQ {
	items := make([]*State, 10)
	return MinPQ{items: items}
}

func (q *MinPQ) isEmpty() bool {
	return q.n == 0
}

func (q *MinPQ) greater(i, j int) bool {
	return q.items[i].estimate > q.items[j].estimate
}

func (q *MinPQ) exch(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *MinPQ) resize(capacity int) {
	cpy := make([]*State, capacity)
	copy(cpy, q.items)
	q.items = cpy
}

func (q *MinPQ) swim(k int) {
	for k > 1 && q.greater(k/2, k) {
		q.exch(k, k/2)
		k = k / 2
	}
}

func (q *MinPQ) insert(s *State) {
	if q.n == len(q.items)-1 {
		q.resize(2 * len(q.items))
	}
	q.n++
	q.items[q.n] = s
	q.swim(q.n)
}

func (q *MinPQ) sink(k int) {
	for 2*k <= q.n {
		j := 2 * k
		if j < q.n && q.greater(j, j+1) {
			j++
		}
		if !q.greater(k, j) {
			break
		}
		q.exch(k, j)
		k = j
	}
}

func (q *MinPQ) delMin() *State {
	min := q.items[1]
	q.exch(1, q.n)
	q.n--
	q.sink(1)
	q.items[q.n+1] = nil
	if q.n > 0 && q.n == (len(q.items)-1)/4 {
		q.resize(len(q.items) / 2)
	}
	return min
}

func between(n, min, max int) bool {
	return n >= min && n < max
}

func isValidStep(x, y, m, n int) bool {
	return between(x, 0, m) && between(y, 0, n)
}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	seen := make(map[Step]bool)
	pq := NewMinPQ()
	target := Cell{m - 1, n - 1}
	source := NewState(0, 0, k, 0, target)
	if source.x == target.x && source.y == target.y {
		return 0
	}
	pq.insert(source)
	seen[source.Step] = true
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for !pq.isEmpty() {
		from := pq.delMin()
		for _, dir := range dirs {
			toX, toY := from.x+dir[0], from.y+dir[1]
			if toX == target.x && toY == target.y {
				return from.dist + 1
			}
			if isValidStep(toX, toY, m, n) {
				obstCount := from.obstCount - grid[toX][toY]
				to := NewState(toX, toY, obstCount, from.dist+1, target)
				if obstCount >= 0 && !seen[to.Step] {
					pq.insert(to)
					seen[to.Step] = true
				}
			}
		}
	}
	return -1
}
