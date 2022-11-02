package topkfreq

type Freq struct {
	num   int
	count int
}

type MinPQ struct {
	items []*Freq
	size  int
}

func NewMinPQ() MinPQ {
	items := make([]*Freq, 10)
	return MinPQ{items: items}
}

func (q *MinPQ) isEmpty() bool {
	return q.size == 0
}

func (q *MinPQ) exch(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *MinPQ) greater(i, j int) bool {
	return q.items[i].count > q.items[j].count
}

func (q *MinPQ) resize(capacity int) {
	cpy := make([]*Freq, capacity)
	copy(cpy, q.items)
	q.items = cpy
}

func (q *MinPQ) swim(k int) {
	for k > 1 && q.greater(k/2, k) {
		q.exch(k, k/2)
		k = k / 2
	}
}

func (q *MinPQ) insert(f *Freq) {
	if q.size == len(q.items)-1 {
		q.resize(2 * len(q.items))
	}
	q.size++
	q.items[q.size] = f
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

func (q *MinPQ) delMin() *Freq {
	item := q.items[1]
	q.exch(1, q.size)
	q.size--
	q.sink(1)
	if q.size == (len(q.items)-1)/4 {
		q.resize(len(q.items) / 2)
	}
	return item
}

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	}
	pq := NewMinPQ()
	for num, freq := range freqs {
		pq.insert(&Freq{num, freq})
		if pq.size > k {
			pq.delMin()
		}
	}
	var result []int
	for i := 1; i <= pq.size; i++ {
		result = append(result, pq.items[i].num)
	}
	return result
}
