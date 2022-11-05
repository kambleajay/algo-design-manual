package furthestbuilding

type MinPQ struct {
	items []int
	size  int
}

func NewMinPQ() MinPQ {
	items := make([]int, 10)
	return MinPQ{items: items}
}

func (q *MinPQ) resize(capacity int) {
	tmp := make([]int, capacity)
	copy(tmp, q.items)
	q.items = tmp
}

func (q *MinPQ) isEmpty() bool {
	return q.size == 0
}

func (q *MinPQ) greater(i, j int) bool {
	return q.items[i] > q.items[j]
}

func (q *MinPQ) exch(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *MinPQ) swim(k int) {
	for k > 1 && q.greater(k/2, k) {
		q.exch(k, k/2)
		k = k / 2
	}
}

func (q *MinPQ) insert(n int) {
	if q.size == len(q.items)-1 {
		q.resize(2 * len(q.items))
	}
	q.size++
	q.items[q.size] = n
	q.swim(q.size)
}

func (q *MinPQ) sink(k int) {
	for 2*k <= q.size {
		j := 2 * k
		if j < q.size && q.greater(j, j+1) {
			j++
		}
		if !q.greater(k, j) {
			break
		}
		q.exch(k, j)
		k = j
	}
}

func (q *MinPQ) delMin() int {
	item := q.items[1]
	q.exch(1, q.size)
	q.size--
	q.sink(1)
	if q.size == (len(q.items)-1)/4 {
		q.resize(len(q.items) / 2)
	}
	return item
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	pq := NewMinPQ()
	for i := 0; i < len(heights)-1; i++ {
		h1, h2 := heights[i], heights[i+1]
		climb := h2 - h1
		if climb > 0 {
			pq.insert(climb)
			if pq.size > ladders {
				minClimb := pq.delMin()
				bricks -= minClimb
				if bricks < 0 {
					return i
				}
			}
		}
	}
	return len(heights) - 1
}
