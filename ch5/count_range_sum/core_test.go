package count_range_sum

import "testing"

func TestCountRangeSum(t *testing.T) {
	var tests = []struct {
		input []int
		lower int
		upper int
		want  int
	}{
		{[]int{-2, 5, -1}, -2, 2, 3},
		{[]int{0}, 0, 0, 1},
	}
	for _, test := range tests {
		if got := countRangeSum(test.input, test.lower, test.upper); got != test.want {
			t.Errorf("countRangeSum(%v, %d, %d) was equal to %d, it should have been %d", test.input, test.lower,
				test.upper, got, test.want)
		}
	}
}
