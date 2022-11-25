package nqueens2

type Square struct {
	x, y int
}

type Solution struct {
	queens []int
	count  int
	n      int
}

func NewSolution(n int) Solution {
	queens := make([]int, n+1)
	return Solution{queens: queens, n: n}
}

func (s *Solution) isSolution(k int) bool {
	return k == s.n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sameDiagonal(s1 Square, s2 Square) bool {
	return abs(s1.x-s2.x) == abs(s1.y-s2.y)
}

func sameColumn(s1 Square, s2 Square) bool {
	return s1.y == s2.y
}

func (s *Solution) constructCandidates(k int) []int {
	var candidates []int
	for i := 1; i <= s.n; i++ {
		legalMove := true
		for j := 1; j < k; j++ {
			currQueenSq := Square{k, i}
			prevQueenSq := Square{j, s.queens[j]}
			if sameDiagonal(currQueenSq, prevQueenSq) || sameColumn(currQueenSq, prevQueenSq) {
				legalMove = false
				break
			}
		}
		if legalMove {
			candidates = append(candidates, i)
		}
	}
	return candidates
}

func (s *Solution) backtrack(k int) {
	if s.isSolution(k) {
		s.count++
	} else {
		k++
		candidates := s.constructCandidates(k)
		for _, candidate := range candidates {
			s.queens[k] = candidate
			s.backtrack(k)
		}
	}
}

func totalNQueens(n int) int {
	s := NewSolution(n)
	s.backtrack(0)
	return s.count
}
