package heightchkr

import "testing"

func TestHeightChecker(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{5, 1, 2, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{}, 0},
		{[]int{1}, 0},
		{nil, 0},
		{[]int{8, 1, 4, 7, 6, 4, 1, 2, 2, 7, 5, 5, 4, 8, 5, 7, 4, 5, 2, 8, 5, 2, 8}, 17},
		{[]int{2, 6, 8, 6, 5, 2, 4, 3, 7, 3, 7, 5, 6, 6, 2, 4, 4, 6, 8, 4, 5}, 16},
	}
	for _, test := range tests {
		if got := heightChecker(test.input); got != test.want {
			t.Errorf("heightChecker(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}
