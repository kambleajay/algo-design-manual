package skyline

import "sort"

type Position struct {
	x      int
	height int
}

type MaxPQ struct {
	a    []int
	size int
}

func NewMaxPQ() *MaxPQ {
	a := make([]int, 10)
	return &MaxPQ{a: a}
}

func (pq *MaxPQ) resize(capacity int) {
	tmp := make([]int, capacity)
	copy(tmp, pq.a)
	pq.a = tmp
}

func (pq *MaxPQ) isEmpty() bool {
	return pq.size == 0
}

func (pq *MaxPQ) less(i, j int) bool {
	return pq.a[i] < pq.a[j]
}

func (pq *MaxPQ) exch(i, j int) {
	pq.a[i], pq.a[j] = pq.a[j], pq.a[i]
}

func (pq *MaxPQ) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *MaxPQ) insert(n int) {
	if pq.size == len(pq.a)-1 {
		pq.resize(2 * len(pq.a))
	}
	pq.size++
	pq.a[pq.size] = n
	pq.swim(pq.size)
}

func (pq *MaxPQ) sink(k int) {
	for 2*k <= pq.size {
		j := 2 * k
		if j < pq.size && pq.less(j, j+1) {
			j++
		}
		if !pq.less(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}

func (pq *MaxPQ) delMax() int {
	item := pq.a[1]
	pq.exch(1, pq.size)
	pq.size--
	pq.sink(1)
	if pq.size == (len(pq.a)-1)/4 {
		pq.resize(len(pq.a) / 2)
	}
	return item
}

func (pq *MaxPQ) max() int {
	return pq.a[1]
}

func prevHeight(result [][]int) int {
	last := result[len(result)-1]
	return last[1]
}

func currHeight(live *MaxPQ) int {
	var height int
	if !live.isEmpty() {
		height = live.max()
	}
	return height
}

func getSkyline(buildings [][]int) [][]int {
	var result [][]int
	var edges []Position
	live, past := NewMaxPQ(), NewMaxPQ()
	for _, building := range buildings {
		left, right, height := building[0], building[1], building[2]
		edges = append(edges, Position{left, height}, Position{right, -height})
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].x < edges[j].x
	})
	var i int
	for i < len(edges) {
		currX := edges[i].x
		for i < len(edges) && edges[i].x == currX {
			height := edges[i].height
			if height > 0 {
				live.insert(height)
			} else {
				past.insert(-height)
			}
			i++
		}
		for !past.isEmpty() && live.max() == past.max() {
			live.delMax()
			past.delMax()
		}
		currHeight := currHeight(live)
		if len(result) == 0 || prevHeight(result) != currHeight {
			result = append(result, []int{currX, currHeight})
		}
	}
	return result
}
