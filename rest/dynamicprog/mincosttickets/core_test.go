package mincosttickets

import "testing"

func TestMinCostTickets(t *testing.T) {
	var tests = []struct {
		days  []int
		costs []int
		want  int
	}{
		{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
	}
	for _, test := range tests {
		if got := mincostTickets(test.days, test.costs); got != test.want {
			t.Errorf("mincostTickets(%v, %v) = %d, want %d", test.days, test.costs, got, test.want)
		}
	}
}
