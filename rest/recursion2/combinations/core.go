package combinations

type Solution struct {
	n, size int
	a       []int
	result  [][]int
}

func NewSolution(n, size int) Solution {
	a := make([]int, size+1)
	return Solution{n: n, size: size, a: a}
}

func (s *Solution) isSolution(k int) bool {
	return k == s.size
}

func (s *Solution) processSolution() {
	tmp := make([]int, s.size)
	copy(tmp, s.a[1:])
	s.result = append(s.result, tmp)
}

func (s *Solution) constructCandidates(k int) []int {
	var candidates []int
	var start int
	if k == 1 {
		start = 1
	} else {
		start = s.a[k-1] + 1
	}
	for i := start; i <= s.n; i++ {
		candidates = append(candidates, i)
	}
	return candidates
}

func (s *Solution) backtrack(k int) {
	if s.isSolution(k) {
		s.processSolution()
	} else {
		k++
		for _, candidate := range s.constructCandidates(k) {
			s.a[k] = candidate
			s.backtrack(k)
		}
	}
}

func combine(n int, k int) [][]int {
	s := NewSolution(n, k)
	s.backtrack(0)
	return s.result
}
