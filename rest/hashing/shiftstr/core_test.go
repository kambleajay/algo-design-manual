package shiftstr

import (
	"algo/testing/utils"
	"testing"
)

func TestHash(t *testing.T) {
	var tests = []struct {
		str  string
		want string
	}{
		{"xyz", "abc"},
		{"bcd", "abc"},
		{"abc", "abc"},
		{"ba", "az"},
	}
	for _, test := range tests[3:] {
		if got := hash(test.str); got != test.want {
			t.Errorf("hash(%s) = %s, want %s", test.str, got, test.want)
		}
	}
}

func TestGroupStrings(t *testing.T) {
	var tests = []struct {
		strs []string
		want [][]string
	}{
		{[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}, [][]string{{"acef"}, {"a", "z"}, {"abc", "bcd", "xyz"}, {"az", "ba"}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, test := range tests {
		if got := groupStrings(test.strs); !utils.Equals(got, test.want) {
			t.Errorf("groupStrings(%v) = %v, want %v", test.strs, got, test.want)
		}
	}
}
