package shortpathobsteli

type State struct {
	x         int
	y         int
	obstCount int
}

type StateWithDist struct {
	State
	dist int
}

type Queue struct {
	items []StateWithDist
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(s StateWithDist) {
	q.items = append(q.items, s)
}

func (q *Queue) dequeue() StateWithDist {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func between(n, min, max int) bool {
	return n >= min && n < max
}

func toState(from StateWithDist, x, y int, grid [][]int) StateWithDist {
	s := StateWithDist{}
	s.x = x
	s.y = y
	s.dist = from.dist + 1
	s.obstCount = from.obstCount + grid[x][y]
	return s
}

func isLowerRight(x, y, m, n int) bool {
	return x == (m-1) && y == (n-1)
}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	seen := make(map[State]bool)
	q := Queue{}
	if isLowerRight(0, 0, m, n) {
		return 0
	}
	source := StateWithDist{State{0, 0, 0}, 0}
	q.enqueue(source)
	seen[source.State] = true
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for !q.isEmpty() {
		v := q.dequeue()
		for _, dir := range dirs {
			wx, wy := v.x+dir[0], v.y+dir[1]
			if between(wx, 0, m) && between(wy, 0, n) {
				nextState := toState(v, wx, wy, grid)
				if nextState.obstCount <= k {
					if isLowerRight(nextState.x, nextState.y, m, n) {
						return nextState.dist
					} else if !seen[nextState.State] {
						q.enqueue(nextState)
						seen[nextState.State] = true
					}
				}
			}
		}
	}
	return -1
}
