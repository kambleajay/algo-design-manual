package longcomsubs

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	var tests = []struct {
		text1, text2 string
		want         int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"abcdef", "f", 1},
		{"ezupkr", "ubmrapg", 2},
	}
	for _, test := range tests[4:] {
		if got := longestCommonSubsequence(test.text1, test.text2); got != test.want {
			t.Errorf("longestCommonSubsequence(%s, %s) = %d, want %d", test.text1, test.text2, got, test.want)
		}
	}
}
