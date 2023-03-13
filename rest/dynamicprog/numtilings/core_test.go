package numtilings

import "testing"

func TestNumTilings(t *testing.T) {
	var tests = []struct {
		n, want int
	}{
		{3, 5},
		{1, 1},
		{4, 11},
		{5, 24},
		{30, 312342182},
	}
	for _, test := range tests {
		if got := numTilings(test.n); got != test.want {
			t.Errorf("numTilings(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
