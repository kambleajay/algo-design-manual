package isomor

import "testing"

func TestIsIsomorphic(t *testing.T) {
	var tests = []struct {
		s    string
		t    string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"badc", "baba", false},
	}
	for _, test := range tests {
		if got := isIsomorphic(test.s, test.t); got != test.want {
			t.Errorf("isIsomorphic(%s, %s) = %t, want %t", test.s, test.t, got, test.want)
		}
	}
}
