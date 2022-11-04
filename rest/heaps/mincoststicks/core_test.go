package mincoststicks

import "testing"

func TestConnectSticks(t *testing.T) {
	var tests = []struct {
		sticks []int
		want   int
	}{
		{[]int{2, 4, 3}, 14},
		{[]int{1, 8, 3, 5}, 30},
		{[]int{5}, 0},
	}
	for _, test := range tests {
		if got := connectSticks(test.sticks); got != test.want {
			t.Errorf("connectSticks(%v) = %d, want %d", test.sticks, got, test.want)
		}
	}
}
