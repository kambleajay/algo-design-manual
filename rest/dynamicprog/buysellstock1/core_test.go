package buysellstock1

import "testing"

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		if got := maxProfit(test.prices); got != test.want {
			t.Errorf("maxProfit(%v) = %d, want %d", test.prices, got, test.want)
		}
	}
}
