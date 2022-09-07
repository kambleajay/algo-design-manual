package maxconones

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 0, 0, 1, 1, 0, 1}, 4},
		{[]int{1, 0, 1, 1, 0}, 4},
		{[]int{}, 0},
		{[]int{1, 0, 1, 1, 0, 1}, 4},
	}
	for _, test := range tests {
		if got := findMaxConsecutiveOnes(test.input); got != test.want {
			t.Errorf("findMaxConsecutiveOnes(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}
