package earfullbloom

import "sort"

type State struct {
	plantTime []int
	growTime  []int
	noOfDays  int
}

type Node struct {
	state *State
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}

func (q *Queue) enqueue(s *State) {
	oldTail := q.tail
	q.tail = &Node{state: s}
	if q.isEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
}

func (q *Queue) dequeue() *State {
	item := q.head.state
	q.head = q.head.next
	if q.isEmpty() {
		q.tail = nil
	}
	return item
}

func nextStates(s *State, n int) []*State {
	var states []*State
	for i := 0; i < n; i++ {
		copyOfPlantTimes := make([]int, n)
		copy(copyOfPlantTimes, s.plantTime)
		copyOfGrowTimes := make([]int, n)
		copy(copyOfGrowTimes, s.growTime)
		pt := s.plantTime[i]
		for i := 0; i < n; i++ {
			if copyOfPlantTimes[i] == 0 && copyOfGrowTimes[i] > 0 {
				copyOfGrowTimes[i]--
			}
		}
		if pt > 0 {
			copyOfPlantTimes[i]--
		}
		states = append(states, &State{copyOfPlantTimes, copyOfGrowTimes, s.noOfDays + 1})
	}
	return states
}

func areAllZero(xs []int) bool {
	for _, x := range xs {
		if x != 0 {
			return false
		}
	}
	return true
}

func isFullBloomState(s *State) bool {
	return areAllZero(s.plantTime) && areAllZero(s.growTime)
}

func earliestFullBloom1(plantTime []int, growTime []int) int {
	n := len(plantTime)
	q := Queue{}
	source := &State{plantTime, growTime, 0}
	q.enqueue(source)
	for !q.isEmpty() {
		currState := q.dequeue()
		for _, nextState := range nextStates(currState, n) {
			if isFullBloomState(nextState) {
				return currState.noOfDays + 1
			}
			q.enqueue(nextState)
		}
	}
	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Tuple struct {
	growTime int
	index    int
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	indices := make([]Tuple, len(growTime))
	for i := 0; i < len(growTime); i++ {
		indices[i] = Tuple{growTime[i], i}
	}
	sort.Slice(indices, func(i, j int) bool {
		return indices[i].growTime > indices[j].growTime
	})
	var result, currPlantTime int
	for _, t := range indices {
		currPlantTime += plantTime[t.index]
		result = max(result, currPlantTime+growTime[t.index])

	}
	return result
}
