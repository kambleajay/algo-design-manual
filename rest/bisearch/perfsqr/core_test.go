package perfsqr

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	var tests = []struct {
		n    int
		want bool
	}{
		{16, true},
		{14, false},
		{1, true},
	}
	for _, test := range tests {
		if got := isPerfectSquare(test.n); got != test.want {
			t.Errorf("isPerfectSquare(%d) = %t, want %t", test.n, got, test.want)
		}
	}
}
