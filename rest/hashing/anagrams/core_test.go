package anagrams

import (
	"algo/testing/utils"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	var tests = []struct {
		strs []string
		want [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, test := range tests {
		if got := groupAnagrams(test.strs); !utils.Equals(got, test.want) {
			t.Errorf("groupAnagrams(%v) = %v, want %v", test.strs, got, test.want)
		}
	}
}
