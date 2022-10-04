package happynum

import "testing"

func TestIsHappy(t *testing.T) {
	var tests = []struct {
		num  int
		want bool
	}{
		{19, true},
		{2, false},
		{1, true},
		{7, true},
		{116, false},
	}
	for _, test := range tests {
		if got := isHappy(test.num); got != test.want {
			t.Errorf("isHappy(%d) = %t, want %t", test.num, got, test.want)
		}
	}
}
