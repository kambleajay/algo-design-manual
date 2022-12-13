package maxscore

import "testing"

func TestMaximumScore(t *testing.T) {
	var tests = []struct {
		nums        []int
		multipliers []int
		want        int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}, 14},
		{[]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}, 102},
	}
	for _, test := range tests {
		if got := maximumScore(test.nums, test.multipliers); got != test.want {
			t.Errorf("maximumScore(%v, %v) = %d, want %d", test.nums, test.multipliers, got, test.want)
		}
	}
}
