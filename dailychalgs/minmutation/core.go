package minmutation

var geneChars = []rune{'A', 'C', 'G', 'T'}

type State struct {
	gene      string
	mutations int
}

type Queue struct {
	items []State
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) enqueue(s State) {
	q.items = append(q.items, s)
}

func (q *Queue) dequeue() State {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func replaceCharAt(s string, i int, r rune) string {
	return s[:i] + string(r) + s[i+1:]
}

func nextStates(s State, end string) []State {
	var states []State
	for i := 0; i < 8; i++ {
		for _, gc := range geneChars {
			if rune(s.gene[i]) != gc {
				newGene := replaceCharAt(s.gene, i, gc)
				newState := State{newGene, s.mutations + 1}
				states = append(states, newState)
			}
		}
	}
	return states
}

func minMutation(start string, end string, bank []string) int {
	bankm := make(map[string]bool)
	for _, gene := range bank {
		bankm[gene] = true
	}
	if !bankm[end] {
		return -1
	}
	q := Queue{}
	source := State{start, 0}
	q.enqueue(source)
	for !q.isEmpty() {
		from := q.dequeue()
		for _, to := range nextStates(from, end) {
			if bankm[to.gene] {
				if to.gene == end {
					return to.mutations
				}
				q.enqueue(to)
			}
		}
	}
	return -1
}
