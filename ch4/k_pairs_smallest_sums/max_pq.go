package k_pairs_smallest_sums

type MaxPQ struct {
	pq [][]int
	N  int
}

func NewMaxPQ() *MaxPQ {
	queue := MaxPQ{}
	queue.pq = make([][]int, 10)
	return &queue
}

func (queue *MaxPQ) isEmpty() bool {
	return queue.N == 0
}

func (queue *MaxPQ) size() int {
	return queue.N
}

func (queue *MaxPQ) less(i, j int) bool {
	i1, i2 := queue.pq[i][0], queue.pq[i][1]
	j1, j2 := queue.pq[j][0], queue.pq[j][1]
	return (i1 + i2) < (j1 + j2)
}

func (queue *MaxPQ) exch(i, j int) {
	queue.pq[i], queue.pq[j] = queue.pq[j], queue.pq[i]
}

func (queue *MaxPQ) resize(size int) {
	res := make([][]int, size)
	copy(res, queue.pq)
	queue.pq = res
}

func (queue *MaxPQ) swim(k int) {
	for k > 1 && queue.less(k/2, k) {
		queue.exch(k/2, k)
		k /= 2
	}
}

func (queue *MaxPQ) insert(n []int) {
	if queue.N == len(queue.pq)-1 {
		queue.resize(2 * len(queue.pq))
	}
	queue.N++
	queue.pq[queue.N] = n
	queue.swim(queue.N)
}

func (queue *MaxPQ) sink(k int) {
	for 2*k <= queue.N {
		j := 2 * k
		if j < queue.N && queue.less(j, j+1) {
			j++
		}
		if !queue.less(k, j) {
			break
		}
		queue.exch(k, j)
		k = j
	}
}

func (queue *MaxPQ) delMax() []int {
	max := queue.pq[1]
	queue.pq[1], queue.pq[queue.N] = queue.pq[queue.N], queue.pq[1]
	queue.N--
	queue.pq[queue.N+1] = nil
	if queue.N > 0 && queue.N == len(queue.pq)/4 {
		queue.resize(len(queue.pq) / 2)
	}
	queue.sink(1)
	return max
}

func kSmallestPairs2(nums1 []int, nums2 []int, k int) [][]int {
	var answer [][]int
	pq := NewMaxPQ()
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			pq.insert([]int{n1, n2})
			if pq.size() > k {
				pq.delMax()
			}
		}
	}
	for !pq.isEmpty() {
		answer = append([][]int{pq.delMax()}, answer...)
	}
	return answer
}
