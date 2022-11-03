package laststoneweight

type MaxPQ struct {
	items []int
	size  int
}

func NewMaxPQ() MaxPQ {
	items := make([]int, 10)
	return MaxPQ{items: items}
}

func (q *MaxPQ) resize(capacity int) {
	tmp := make([]int, capacity)
	copy(tmp, q.items)
	q.items = tmp
}

func (q *MaxPQ) isEmpty() bool {
	return q.size == 0
}

func (q *MaxPQ) less(i, j int) bool {
	return q.items[i] < q.items[j]
}

func (q *MaxPQ) exch(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *MaxPQ) swim(k int) {
	for k > 1 && q.less(k/2, k) {
		q.exch(k, k/2)
		k = k / 2
	}
}

func (q *MaxPQ) insert(n int) {
	if q.size == len(q.items)-1 {
		q.resize(2 * len(q.items))
	}
	q.size++
	q.items[q.size] = n
	q.swim(q.size)
}

func (q *MaxPQ) sink(k int) {
	for 2*k <= q.size {
		j := 2 * k
		if j < q.size && q.less(j, j+1) {
			j++
		}
		if !q.less(k, j) {
			break
		}
		q.exch(k, j)
		k = j
	}
}

func (q *MaxPQ) delMax() int {
	item := q.items[1]
	q.exch(1, q.size)
	q.size--
	q.sink(1)
	if q.size == (len(q.items)-1)/4 {
		q.resize(len(q.items) / 2)
	}
	return item
}

func (q *MaxPQ) max() int {
	return q.items[1]
}

func lastStoneWeight(stones []int) int {
	pq := NewMaxPQ()
	for _, stone := range stones {
		pq.insert(stone)
	}
	for pq.size > 1 {
		s1 := pq.delMax()
		s2 := pq.delMax()
		if s1 > s2 {
			pq.insert(s1 - s2)
		}
	}
	if !pq.isEmpty() {
		return pq.max()
	}
	return 0
}
