package revwords

import "testing"

func TestReverseWords(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
	}
	for _, test := range tests {
		if got := reverseWords(test.s); got != test.want {
			t.Errorf("reverseWords(%s) = %s, want %s", test.s, got, test.want)
		}
	}
}
