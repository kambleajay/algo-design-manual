package buysellcooldown

import "testing"

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		prices []int
		want   int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		{[]int{1}, 0},
	}
	for _, test := range tests {
		if got := maxProfit(test.prices); got != test.want {
			t.Errorf("maxProfit(%v) = %d, want %d", test.prices, got, test.want)
		}
	}
}
