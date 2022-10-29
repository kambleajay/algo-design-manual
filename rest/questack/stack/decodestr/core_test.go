package decodestr

import "testing"

func TestDecodeString(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"10[a]", "aaaaaaaaaa"},
	}
	for _, test := range tests {
		if got := decodeString(test.s); got != test.want {
			t.Errorf("decodeString(%s) = %s, want %s", test.s, got, test.want)
		}
	}
}
