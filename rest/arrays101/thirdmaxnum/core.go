package thirdmaxnum

import "math"

type Set struct {
	items map[int]bool
	min   int
	max   int
	n     int
}

func NewSet() *Set {
	return &Set{items: make(map[int]bool), min: math.MaxInt64, max: math.MinInt64}
}

func (s *Set) Add(n int) {
	_, ok := s.items[n]
	if !ok {
		s.items[n] = true
		if n < s.min {
			s.min = n
		}
		if n > s.max {
			s.max = n
		}
		s.n++
	}
}

func (s *Set) Size() int {
	return s.n
}

func (s *Set) Min() int {
	return s.min
}

func (s *Set) Max() int {
	return s.max
}

func (s *Set) Remove(n int) {
	delete(s.items, n)
	s.n--
	min := math.MaxInt64
	for n, _ := range s.items {
		if n < min {
			min = n
		}
	}
	s.min = min
	max := math.MinInt64
	for n, _ := range s.items {
		if n > max {
			max = n
		}
	}
	s.max = max
}

func thirdMax(nums []int) int {
	maximums := NewSet()
	for _, n := range nums {
		maximums.Add(n)
		if maximums.Size() > 3 {
			maximums.Remove(maximums.Min())
		}
	}
	if maximums.Size() == 3 {
		return maximums.Min()
	}
	return maximums.Max()
}
