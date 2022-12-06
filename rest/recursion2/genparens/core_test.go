package genparens

import (
	"algo/testing/utils"
	"testing"
)

func TestGenerateParenthese(t *testing.T) {
	var tests = []struct {
		n    int
		want []string
	}{
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{3, []string{"()()()", "()(())", "(()())", "(())()", "((()))"}},
	}
	for _, test := range tests {
		if got := generateParenthesis(test.n); !utils.ContainsAllStringSlice(got, test.want) {
			t.Errorf("generateParenthesis(%d) = %v, want %v", test.n, got, test.want)
		}
	}
}
