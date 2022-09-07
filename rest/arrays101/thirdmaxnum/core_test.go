package thirdmaxnum

import "testing"

func TestThirdMax(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{12, 3, 8, 9, 12, 12, 7, 8, 12, 4, 3, 8, 1}, 8},
		{[]int{13, 2, 6, 23, 12, 15, 7, 23, 25, 4, 6}, 15},
	}
	for _, test := range tests {
		if got := thirdMax(test.input); got != test.want {
			t.Errorf("thirdMax(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}
