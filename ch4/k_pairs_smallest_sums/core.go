package k_pairs_smallest_sums

type Node struct {
	n1 int
	n2 []int
}

type MinPQ struct {
	nodes []*Node
	N     int
}

func NewMinPQ() *MinPQ {
	queue := MinPQ{}
	queue.nodes = make([]*Node, 10)
	return &queue
}

func (queue *MinPQ) isEmpty() bool {
	return queue.N == 0
}

func (queue *MinPQ) greater(i, j int) bool {
	s1 := queue.nodes[i].n1 + queue.nodes[i].n2[0]
	s2 := queue.nodes[j].n1 + queue.nodes[j].n2[0]
	return s1 > s2
}

func (queue *MinPQ) exch(i, j int) {
	queue.nodes[i], queue.nodes[j] = queue.nodes[j], queue.nodes[i]
}

func (queue *MinPQ) resize(size int) {
	res := make([]*Node, size)
	copy(res, queue.nodes)
	queue.nodes = res
}

func (queue *MinPQ) swim(k int) {
	for k > 1 && queue.greater(k/2, k) {
		queue.exch(k/2, k)
		k /= 2
	}
}

func (queue *MinPQ) insert(n *Node) {
	if queue.N == len(queue.nodes)-1 {
		queue.resize(2 * len(queue.nodes))
	}
	queue.N++
	queue.nodes[queue.N] = n
	queue.swim(queue.N)
}

func (queue *MinPQ) sink(k int) {
	for 2*k <= queue.N {
		j := 2 * k
		if j < queue.N && queue.greater(j, j+1) {
			j++
		}
		if !queue.greater(k, j) {
			break
		}
		queue.exch(k, j)
		k = j
	}
}

func (queue *MinPQ) delMin() *Node {
	min := queue.nodes[1]
	queue.nodes[1], queue.nodes[queue.N] = queue.nodes[queue.N], queue.nodes[1]
	queue.N--
	queue.nodes[queue.N+1] = nil
	if queue.N > 0 && queue.N == len(queue.nodes)-1/4 {
		queue.resize(len(queue.nodes) / 2)
	}
	queue.sink(1)
	return min
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := NewMinPQ()
	for _, n1 := range nums1 {
		pq.insert(&Node{n1, nums2})
	}
	var res [][]int
	for !pq.isEmpty() && k > 0 {
		minNode := pq.delMin()
		res = append(res, []int{minNode.n1, minNode.n2[0]})
		if len(minNode.n2) > 1 {
			pq.insert(&Node{minNode.n1, minNode.n2[1:]})
		}
		k--
	}
	return res
}
