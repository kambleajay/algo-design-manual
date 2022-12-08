package lettercombs

type Solution struct {
	a              [5]byte
	digits         []byte
	result         []string
	n              int
	digitToLetters map[byte][]byte
}

func NewSolution(digits string) Solution {
	digitToLetters := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	return Solution{digits: []byte(digits), n: len(digits), digitToLetters: digitToLetters}
}

func (s *Solution) isSolution(k int) bool {
	return k == s.n
}

func (s *Solution) processSolution(k int) {
	tmp := make([]byte, s.n)
	currSolution := s.a[1 : k+1]
	copy(tmp, currSolution)
	s.result = append(s.result, string(tmp))
}

func (s *Solution) constructCandidates(k int) []byte {
	digit := s.digits[k-1]
	return s.digitToLetters[digit]
}

func (s *Solution) backtrack(k int) {
	if s.isSolution(k) {
		s.processSolution(k)
	} else {
		k++
		for _, c := range s.constructCandidates(k) {
			s.a[k] = c
			s.backtrack(k)
		}
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	solution := NewSolution(digits)
	solution.backtrack(0)
	return solution.result
}
