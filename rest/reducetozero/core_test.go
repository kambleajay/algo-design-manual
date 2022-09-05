package reducetozero

import "testing"

func TestNumberOfSteps(t *testing.T) {
	var tests = []struct {
		num  int
		want int
	}{
		{14, 6},
		{8, 4},
		{123, 12},
		{0, 0},
	}
	for _, test := range tests {
		if got := numberOfSteps(test.num); got != test.want {
			t.Errorf("numberOfSteps(%d) = %d, want %d", test.num, got, test.want)
		}
	}
}
