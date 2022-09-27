package pivotsum

import (
	"testing"
)

func TestPivotSum(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}
	for _, test := range tests {
		if got := pivotIndex(test.input); got != test.want {
			t.Errorf("pivotIndex(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}
