package powxn

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	var tests = []struct {
		b    float64
		n    int
		want float64
	}{
		{2.0, 10, 1024.0},
		{2.1, 3, 9.261},
		{2.0, -2, 0.25},
	}
	for _, test := range tests {
		if got := myPow(test.b, test.n); math.Abs(got-test.want) > 0.0001 {
			t.Errorf("myPow(%g, %d) = %g, want %g", test.b, test.n, got, test.want)
		}
	}
}
