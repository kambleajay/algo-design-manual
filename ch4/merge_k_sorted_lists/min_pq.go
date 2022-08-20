package merge_k_sorted_lists

type MinPQ struct {
	nodes []*ListNode
	N     int
}

func NewMinPQ() *MinPQ {
	queue := MinPQ{}
	queue.nodes = make([]*ListNode, 10)
	return &queue
}

func (queue *MinPQ) isEmpty() bool {
	return queue.N == 0
}

func (queue *MinPQ) greater(i, j int) bool {
	return queue.nodes[i].Val > queue.nodes[j].Val
}

func (queue *MinPQ) exch(i, j int) {
	queue.nodes[i], queue.nodes[j] = queue.nodes[j], queue.nodes[i]
}

func (queue *MinPQ) resize(size int) {
	res := make([]*ListNode, size)
	copy(res, queue.nodes)
	queue.nodes = res
}

func (queue *MinPQ) swim(k int) {
	for k > 1 && queue.greater(k/2, k) {
		queue.exch(k/2, k)
		k /= 2
	}
}

func (queue *MinPQ) insert(n *ListNode) {
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

func (queue *MinPQ) delMin() *ListNode {
	min := queue.nodes[1]
	queue.nodes[1], queue.nodes[queue.N] = queue.nodes[queue.N], queue.nodes[1]
	queue.N--
	queue.nodes[queue.N+1] = nil
	if queue.N > 0 && queue.N == (len(queue.nodes))-1/4 {
		queue.resize(len(queue.nodes) / 2)
	}
	queue.sink(1)
	return min
}

func mergeKLists2(lists []*ListNode) *ListNode {
	queue := NewMinPQ()

	for _, list := range lists {
		if list != nil {
			queue.insert(list)
		}
	}

	head := &ListNode{}
	tail := head

	for !queue.isEmpty() {
		minNode := queue.delMin()
		tail.Next = minNode
		tail = tail.Next
		if minNode.Next != nil {
			queue.insert(minNode.Next)
		}
	}
	return head.Next
}
