package mindifficulty

import "testing"

func TestMinDifficulty(t *testing.T) {
	var tests = []struct {
		jobDifficulty []int
		d             int
		want          int
	}{
		{[]int{6, 5, 4, 3, 2, 1}, 2, 7},
		{[]int{9, 9, 9}, 4, -1},
		{[]int{1, 1, 1}, 3, 3},
	}
	for _, test := range tests {
		if got := minDifficulty(test.jobDifficulty, test.d); got != test.want {
			t.Errorf("minDifficulty(%v, %d) = %d, want %d", test.jobDifficulty, test.d, got, test.want)
		}
	}
}
