package sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	var tests = []struct {
		num  int
		want int
	}{
		{4, 2},
		{8, 2},
	}
	for _, test := range tests {
		if got := mySqrt(test.num); got != test.want {
			t.Errorf("mySqrt(%d) = %d, want %d", test.num, got, test.want)
		}
	}
}
