package laststoneweight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	var tests = []struct {
		stones []int
		want   int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1}, 1},
		{[]int{2, 2}, 0},
	}
	for _, test := range tests {
		if got := lastStoneWeight(test.stones); got != test.want {
			t.Errorf("lastStoneWeight(%v) = %d, want %d", test.stones, got, test.want)
		}
	}
}
