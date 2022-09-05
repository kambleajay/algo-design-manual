package rects

import (
	"math/rand"
	"time"
)

type Solution struct {
	rng           *rand.Rand
	totalWeight   int
	weightedRects []*WeightedRect
}

type WeightedRect struct {
	min, max int
	rect     []int
}

func areaOfRect(rect []int) int {
	x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
	return (x2 - x1 + 1) * (y2 - y1 + 1)
}

func Constructor(rects [][]int) Solution {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	var weightedRects []*WeightedRect
	var sum int
	for i, r := range rects {
		areaR := areaOfRect(r)
		sum += areaR
		if i == 0 {
			weightedRects = append(weightedRects, &WeightedRect{0, sum - 1, r})
		} else {
			weightedRects = append(weightedRects, &WeightedRect{sum - areaR, sum - 1, r})
		}
	}
	return Solution{rng, sum, weightedRects}
}

func (s *Solution) pickRandomRect() []int {
	n := s.rng.Intn(s.totalWeight)
	lo, hi := 0, len(s.weightedRects)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		midRect := s.weightedRects[mid]
		if n >= midRect.min && n <= midRect.max {
			return midRect.rect
		} else if n < midRect.min {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return nil
}

func (s *Solution) Pick() []int {
	randomRect := s.pickRandomRect()
	x1, y1, x2, y2 := randomRect[0], randomRect[1], randomRect[2], randomRect[3]
	randomX := s.rng.Intn(x2-x1+1) + x1
	randomY := s.rng.Intn(y2-y1+1) + y1
	return []int{randomX, randomY}
}
