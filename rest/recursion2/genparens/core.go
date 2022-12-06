package genparens

const (
	Open  = '('
	Close = ')'
)

type Solution struct {
	a      []byte
	result []string
}

func NewSolution(n int) Solution {
	a := make([]byte, (2*n)+1)
	return Solution{a: a}
}

func (s *Solution) isSolution(openCount, closeCount int) bool {
	return openCount == 0 && closeCount == 0
}

func (s *Solution) processSolution(k int) {
	currSolution := string(s.a[1 : k+1])
	s.result = append(s.result, currSolution)
}

func (s *Solution) constructCandidates(k, openCount, closeCount int) []byte {
	var candidates []byte
	if openCount > 0 {
		candidates = append(candidates, Open)
	}
	if openCount < closeCount {
		candidates = append(candidates, Close)
	}
	return candidates
}

func (s *Solution) backtrack(k, openCount, closeCount int) {
	if s.isSolution(openCount, closeCount) {
		s.processSolution(k)
	} else {
		k++
		for _, candidate := range s.constructCandidates(k, openCount, closeCount) {
			s.a[k] = candidate
			if candidate == Open {
				s.backtrack(k, openCount-1, closeCount)
			} else {
				s.backtrack(k, openCount, closeCount-1)
			}
		}
	}
}

func generateParenthesis(n int) []string {
	s := NewSolution(n)
	s.backtrack(0, n, n)
	return s.result
}
