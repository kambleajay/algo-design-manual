package permutations

type Solution struct {
	nums   []int
	a      []int
	result [][]int
	n      int
}

func NewSolution(nums []int) Solution {
	a := make([]int, len(nums)+1)
	return Solution{nums: nums, a: a, n: len(nums)}
}

func (s *Solution) isSolution(k int) bool {
	return k == s.n
}

func (s *Solution) processSolution(k int) {
	tmp := make([]int, s.n)
	currSolution := s.a[1 : k+1]
	copy(tmp, currSolution)
	s.result = append(s.result, tmp)
}

func (s *Solution) inCurrPermutation(k, num int) bool {
	for i := 1; i < k; i++ {
		if s.a[i] == num {
			return true
		}
	}
	return false
}

func (s *Solution) constructCandidates(k int) []int {
	var result []int
	for i := 0; i < s.n; i++ {
		if !s.inCurrPermutation(k, s.nums[i]) {
			result = append(result, s.nums[i])
		}
	}
	return result
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

func permute(nums []int) [][]int {
	solution := NewSolution(nums)
	solution.backtrack(0)
	return solution.result
}
