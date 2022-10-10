package topkfreq

type Freq struct {
	num  int
	freq int
}

type MinPQ struct {
	freqs []*Freq
	size  int
}

func NewMinPQ() *MinPQ {
	a := make([]*Freq, 1)
	return &MinPQ{freqs: a}
}

func (q *MinPQ) isEmpty() bool {
	return q.size == 0
}

func (q *MinPQ) resize(capacity int) {
	cp := make([]*Freq, capacity)
	copy(cp[1:], q.freqs[1:q.size+1])
	q.freqs = cp
}

func (q *MinPQ) greater(i, j int) bool {
	return q.freqs[i].freq > q.freqs[j].freq
}

func (q *MinPQ) exch(i, j int) {
	q.freqs[i], q.freqs[j] = q.freqs[j], q.freqs[i]
}

func (q *MinPQ) swim(k int) {
	for k > 1 && q.greater(k/2, k) {
		q.exch(k, k/2)
		k = k / 2
	}
}

func (q *MinPQ) insert(f *Freq) {
	if q.size == len(q.freqs)-1 {
		q.resize(2 * len(q.freqs))
	}
	q.size++
	q.freqs[q.size] = f
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
	min := q.freqs[1]
	q.exch(1, q.size)
	q.size--
	q.sink(1)
	q.freqs[q.size+1] = nil
	if q.size > 0 && q.size == (len(q.freqs)-1)/4 {
		q.resize(len(q.freqs) / 2)
	}
	return min.num
}

func (q *MinPQ) elems() []*Freq {
	return q.freqs[1 : q.size+1]
}

func topKFrequent1(nums []int, k int) []int {
	freqs := make(map[int]int)
	for _, n := range nums {
		freqs[n]++
	}
	minPQ := NewMinPQ()
	for num, count := range freqs {
		minPQ.insert(&Freq{num, count})
		if minPQ.size > k {
			minPQ.delMin()
		}
	}
	var result []int
	for _, freq := range minPQ.elems() {
		result = append(result, freq.num)
	}
	return result
}

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)
	for _, n := range nums {
		freqs[n]++
	}
	freqBuckets := make([][]int, len(nums)+1)
	for num, freq := range freqs {
		freqBuckets[freq] = append(freqBuckets[freq], num)
	}
	var result []int
	for j := len(freqBuckets) - 1; j > 0; j-- {
		for _, i := range freqBuckets[j] {
			result = append(result, i)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}
