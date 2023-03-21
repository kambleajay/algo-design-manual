package isinterleave

import "testing"

func TestIsInterleave(t *testing.T) {
	var tests = []struct {
		s1, s2, s3 string
		want       bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"", "", "a", false},
	}
	for _, test := range tests {
		if got := isInterleave(test.s1, test.s2, test.s3); got != test.want {
			t.Errorf("isInterleave(%s, %s, %s) = %t, want %t", test.s1, test.s2, test.s3, got, test.want)
		}
	}
}
