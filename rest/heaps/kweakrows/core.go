package kweakrows

type Row struct {
	soldierCount int
	index        int
}

type MaxPQ struct {
	items []*Row
	size  int
}

func NewMaxPQ() MaxPQ {
	items := make([]*Row, 10)
	return MaxPQ{items: items}
}

func (q *MaxPQ) resize(capacity int) {
	tmp := make([]*Row, capacity)
	copy(tmp, q.items)
	q.items = tmp
}

func (q *MaxPQ) isEmpty() bool {
	return q.size == 0
}

func countOfSoldiers(folks []int) int {
	var count int
	for ; count < len(folks) && folks[count] == 1; count++ {
	}
	return count
}

func (q *MaxPQ) less(i, j int) bool {
	iSoldierCount := q.items[i].soldierCount
	jSoldierCount := q.items[j].soldierCount
	lessSoldiers := iSoldierCount < jSoldierCount
	equalSoldiers := iSoldierCount == jSoldierCount
	lessIndex := q.items[i].index < q.items[j].index
	return lessSoldiers || (equalSoldiers && lessIndex)
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

func (q *MaxPQ) insert(row *Row) {
	if q.size == len(q.items)-1 {
		q.resize(2 * len(q.items))
	}
	q.size++
	q.items[q.size] = row
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

func (q *MaxPQ) delMax() *Row {
	item := q.items[1]
	q.exch(1, q.size)
	q.items[q.size] = nil
	q.size--
	q.sink(1)
	if q.size == (len(q.items)-1)/4 {
		q.resize(len(q.items) / 2)
	}
	return item
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func kWeakestRows(mat [][]int, k int) []int {
	pq := NewMaxPQ()
	for i, row := range mat {
		pq.insert(&Row{countOfSoldiers(row), i})
		if pq.size > k {
			pq.delMax()
		}
	}
	var result []int
	for !pq.isEmpty() {
		result = append(result, pq.delMax().index)
	}
	reverse(result)
	return result
}
