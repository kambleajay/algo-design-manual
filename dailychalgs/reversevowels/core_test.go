package reversevowels

import "testing"

func TestReverseVowels(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aeiou", "uoiea"},
		{"xyzdc", "xyzdc"},
	}
	for _, test := range tests {
		if got := reverseVowels(test.s); got != test.want {
			t.Errorf("reverseVowels(%s) = %s, want %s", test.s, got, test.want)
		}
	}
}
