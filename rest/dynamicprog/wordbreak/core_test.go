package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	var tests = []struct {
		s     string
		words []string
		want  bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, test := range tests {
		if got := wordBreak(test.s, test.words); got != test.want {
			t.Errorf("wordBreak(%s, %v) = %t, want %t", test.s, test.words, got, test.want)
		}
	}
}
