package opnlck

import (
	"testing"
)

func TestOpenLock(t *testing.T) {
	var tests = []struct {
		deadends []string
		target   string
		want     int
	}{
		{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
		{[]string{"8888"}, "0009", 1},
		{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
	}
	for i, test := range tests {
		if got := openLock(test.deadends, test.target); got != test.want {
			t.Errorf("[%d] openLock(%v, %s) = %d, want %d", i, test.deadends, test.target, got, test.want)
		}
	}
}
