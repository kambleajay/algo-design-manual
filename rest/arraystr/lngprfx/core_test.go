package lngprfx

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"solo"}, "solo"},
		{[]string{"a"}, "a"},
	}
	for _, test := range tests {
		if got := longestCommonPrefix(test.strs); got != test.want {
			t.Errorf("longestCommonPrefix(%v) = %s, want %s", test.strs, got, test.want)
		}
	}
}
