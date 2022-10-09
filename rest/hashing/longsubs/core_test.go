package longsubs

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"h", 1},
		{"dvdf", 3},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, test := range tests {
		if got := lengthOfLongestSubstring(test.s); got != test.want {
			t.Errorf("lengthOfLongestSubstring(%s) = %d, want %d", test.s, got, test.want)
		}
	}
}
