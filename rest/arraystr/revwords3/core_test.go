package revwords3

import "testing"

func TestReverseWords(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"God Ding", "doG gniD"},
		{"", ""},
	}
	for _, test := range tests {
		if got := reverseWords(test.s); got != test.want {
			t.Errorf("reverseWords(%s) = %s, want %s", test.s, got, test.want)
		}
	}
}
