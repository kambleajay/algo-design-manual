package division

import "testing"

func TestDivide(t *testing.T) {
	var tests = []struct {
		dividend float64
		divisor  float64
		want     float64
	}{
		{27, 9, 3},
		{100, 10, 10},
		{46, 5, 9},
	}
	for _, test := range tests {
		if got := Divide2(test.dividend, test.divisor); got != test.want {
			t.Errorf("%f/%f != %f (%f)\n", test.dividend, test.divisor, test.want, got)
		}
	}
}
