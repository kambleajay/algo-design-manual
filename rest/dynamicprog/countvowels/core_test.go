package countvowels

import "testing"

func TestCountVowelPermutation(t *testing.T) {
	var tests = []struct {
		n, want int
	}{
		{1, 5},
		{2, 10},
		{5, 68},
		{144, 18208803},
	}
	for _, test := range tests {
		if got := countVowelPermutation(test.n); got != test.want {
			t.Errorf("countVowelPermutation(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
