package dicerolls

import "testing"

func TestNumOfRollsToTarget(t *testing.T) {
	var tests = []struct {
		n, k, target int
		want         int
	}{
		{1, 6, 3, 1},
		{2, 6, 7, 6},
		{30, 30, 500, 222616187},
	}
	for _, test := range tests {
		if got := numRollsToTarget(test.n, test.k, test.target); got != test.want {
			t.Errorf("numRollsToTarget(%d, %d, %d) = %d, want %d", test.n, test.k, test.target, got, test.want)
		}
	}
}
