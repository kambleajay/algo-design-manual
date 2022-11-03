package meetrooms2

type Meeting struct {
	start int
	end   int
}

type MinPQ struct {
	items         []*Meeting
	size          int
	greaterHelper func(*Meeting, *Meeting) bool
}

func NewMinPQ(greaterFn func(*Meeting, *Meeting) bool) MinPQ {
	items := make([]*Meeting, 10)
	return MinPQ{items: items, greaterHelper: greaterFn}
}

func (q *MinPQ) resize(capacity int) {
	tmp := make([]*Meeting, capacity)
	copy(tmp, q.items)
	q.items = tmp
}

func (q *MinPQ) isEmpty() bool {
	return q.size == 0
}

func (q *MinPQ) greater(i, j int) bool {
	return q.greaterHelper(q.items[i], q.items[j])
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

func (q *MinPQ) insert(m *Meeting) {
	if q.size == len(q.items)-1 {
		q.resize(2 * len(q.items))
	}
	q.size++
	q.items[q.size] = m
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

func (q *MinPQ) delMin() *Meeting {
	item := q.items[1]
	q.exch(1, q.size)
	q.size--
	q.sink(1)
	if q.size == (len(q.items)-1)/4 {
		q.resize(len(q.items) / 2)
	}
	return item
}

func (q *MinPQ) min() *Meeting {
	return q.items[1]
}

func greaterByStartTime(m1, m2 *Meeting) bool {
	return m1.start > m2.start
}

func greaterByEndTime(m1, m2 *Meeting) bool {
	return m1.end > m2.end
}

func minMeetingRooms(intervals [][]int) int {
	startTimesPQ := NewMinPQ(greaterByStartTime)
	endTimesPQ := NewMinPQ(greaterByEndTime)
	for _, meeting := range intervals {
		start, end := meeting[0], meeting[1]
		startTimesPQ.insert(&Meeting{start, end})
	}
	endTimesPQ.insert(startTimesPQ.delMin())
	for !startTimesPQ.isEmpty() {
		meeting := startTimesPQ.delMin()
		if meeting.start >= endTimesPQ.min().end {
			endTimesPQ.delMin()
		}
		endTimesPQ.insert(meeting)
	}
	return endTimesPQ.size
}
