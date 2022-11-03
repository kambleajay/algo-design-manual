package kclosepoints

import "math"

type MaxPQ struct {
	items [][]int
	size  int
}

func NewMaxPQ() MaxPQ {
	items := make([][]int, 10)
	return MaxPQ{items: items}
}

func (q *MaxPQ) resize(capacity int) {
	tmp := make([][]int, capacity)
	copy(tmp, q.items)
	q.items = tmp
}

func (q *MaxPQ) isEmpty() bool {
	return q.size == 0
}

func distToOrigin(p []int) float64 {
	xsquared := float64(p[0] * p[0])
	ysquared := float64(p[1] * p[1])
	return math.Sqrt(xsquared + ysquared)
}

func (q *MaxPQ) less(i, j int) bool {
	return distToOrigin(q.items[i]) < distToOrigin(q.items[j])
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

func (q *MaxPQ) insert(p []int) {
	if q.size == len(q.items)-1 {
		q.resize(2 * len(q.items))
	}
	q.size++
	q.items[q.size] = p
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

func (q *MaxPQ) delMax() []int {
	item := q.items[1]
	q.exch(1, q.size)
	q.size--
	q.sink(1)
	if q.size == (len(q.items)-1)/4 {
		q.resize(len(q.items) / 2)
	}
	return item
}

func kClosest(points [][]int, k int) [][]int {
	pq := NewMaxPQ()
	for _, point := range points {
		pq.insert(point)
		if pq.size > k {
			pq.delMax()
		}
	}
	var result [][]int
	for !pq.isEmpty() {
		result = append(result, pq.delMax())
	}
	return result
}
