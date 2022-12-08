package lettercombs

import (
	"algo/testing/utils"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	var tests = []struct {
		digits string
		want   []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", nil},
		{"2", []string{"a", "b", "c"}},
		{"27", []string{"ap", "aq", "ar", "as", "bp", "bq", "br", "bs", "cp", "cq", "cr", "cs"}},
	}
	for _, test := range tests {
		if got := letterCombinations(test.digits); !utils.ContainsAllStringSlice(got, test.want) {
			t.Errorf("letterCombinations(%s) = %v, want %v", test.digits, got, test.want)
		}
	}
}
