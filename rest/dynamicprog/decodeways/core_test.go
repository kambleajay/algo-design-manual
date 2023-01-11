package decodeways

import "testing"

func TestNumDecodings(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"11106", 2},
		{"12", 2},
		{"226", 3},
		{"06", 0},
		{"99", 1},
		{"0", 0},
		{"10", 1},
		{"123123", 9},
	}
	for _, test := range tests {
		if got := numDecodings(test.s); got != test.want {
			t.Errorf("numDecodings(%s) = %d, want %d", test.s, got, test.want)
		}
	}
}
