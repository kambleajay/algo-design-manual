package mincostclimbstairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	var tests = []struct {
		cost []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{[]int{5, 10}, 5},
	}
	for _, test := range tests {
		if got := minCostClimbingStairs(test.cost); got != test.want {
			t.Errorf("minCostClimbingStairs(%v) = %d, want %d", test.cost, got, test.want)
		}
	}
}
