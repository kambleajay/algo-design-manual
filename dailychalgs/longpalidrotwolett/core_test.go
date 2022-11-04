package longpalidrotwolett

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		words []string
		want  int
	}{
		{[]string{"lc", "cl", "gg"}, 6},
		{[]string{"ab", "ty", "yt", "lc", "cl", "ab"}, 8},
		{[]string{"cc", "ll", "xx"}, 2},
		{[]string{"io", "io"}, 0},
		{[]string{"ab", "xx", "yy", "zz", "ba"}, 6},
		{[]string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}, 22},
	}
	for _, test := range tests {
		if got := longestPalindrome(test.words); got != test.want {
			t.Errorf("longestPalindrome(%v) = %d, want %d", test.words, got, test.want)
		}
	}
}
