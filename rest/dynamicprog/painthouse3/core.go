package painthouse3

const MaxCost = int(1e7)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type State struct {
	houseIndex        int
	noOfNeighborhoods int
	prevHouseColor    int
}

type Solution struct {
	houses []int
	cost   [][]int
	m      int
	n      int
	target int
	memo   map[State]int
}

func NewSolution(houses []int, cost [][]int, m, n, target int) Solution {
	memo := make(map[State]int)
	return Solution{houses, cost, m, n, target, memo}
}

func isAlreadyPainted(color int) bool {
	return color != 0
}

func countNoOfNeighborHoods(currHouseColor int, prevHouseColor int, noOfNeighborHoods int) int {
	if currHouseColor != prevHouseColor {
		return noOfNeighborHoods + 1
	}
	return noOfNeighborHoods
}

func (s *Solution) dpfn(i int, prevHouseColor int, noOfNeighborHoods int) int {
	if i == s.m {
		if noOfNeighborHoods == s.target {
			return 0
		} else {
			return MaxCost
		}
	}
	if noOfNeighborHoods > s.target {
		return MaxCost
	}
	state := State{i, noOfNeighborHoods, prevHouseColor}
	answer, ok := s.memo[state]
	if ok {
		return answer
	}
	currHouseColor := s.houses[i]
	minCost := MaxCost
	if isAlreadyPainted(currHouseColor) {
		minCost = s.dpfn(i+1, currHouseColor, countNoOfNeighborHoods(currHouseColor, prevHouseColor, noOfNeighborHoods))
	} else {
		for color := 1; color <= s.n; color++ {
			colorCost := s.cost[i][color-1]
			costWithCurrColor := colorCost + s.dpfn(i+1, color, countNoOfNeighborHoods(color, prevHouseColor, noOfNeighborHoods))
			minCost = min(minCost, costWithCurrColor)
		}
	}
	s.memo[state] = minCost
	return minCost
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	solution := NewSolution(houses, cost, m, n, target)
	answer := solution.dpfn(0, 0, 0)
	if answer == MaxCost {
		return -1
	}
	return answer
}
