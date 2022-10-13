package validparen

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		str  string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"[{()}]", true},
		{"[()]{}", true},
		{"(", false},
		{")", false},
	}
	for _, test := range tests {
		if got := isValid(test.str); got != test.want {
			t.Errorf("isValid(%s) = %t, want %t", test.str, got, test.want)
		}
	}
}
