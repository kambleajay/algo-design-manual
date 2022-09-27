package largtwice

import "testing"

func TestDominantIndex(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{3, 6, 1, 0}, 1},
		{[]int{1, 2, 3, 4}, -1},
		{[]int{1, 1, 1, 2, 4}, 4},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 0},
	}
	for _, test := range tests {
		if got := dominantIndex(test.input); got != test.want {
			t.Errorf("dominantIndex(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}
