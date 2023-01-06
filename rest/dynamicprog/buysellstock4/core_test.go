package buysellstock4

import "testing"

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		k      int
		prices []int
		want   int
	}{
		{2, []int{2, 4, 1}, 2},
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
	}
	for _, test := range tests {
		if got := maxProfit(test.k, test.prices); got != test.want {
			t.Errorf("maxProfit(%d, %v) = %d, want %d", test.k, test.prices, got, test.want)
		}
	}
}
