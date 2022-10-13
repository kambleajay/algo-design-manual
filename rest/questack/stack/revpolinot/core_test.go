package revpolinot

import "testing"

func TestEvalRPN(t *testing.T) {
	var tests = []struct {
		tokens []string
		want   int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, test := range tests {
		if got := evalRPN(test.tokens); got != test.want {
			t.Errorf("evalRPN(%v) = %d, want %d", test.tokens, got, test.want)
		}
	}
}
