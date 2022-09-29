package addbin

import "testing"

func TestAddBinary(t *testing.T) {
	var tests = []struct {
		num1 string
		num2 string
		want string
	}{
		{"1010", "1011", "10101"},
		{"1111", "0010", "10001"},
		{"11", "1", "100"},
		{"1", "0", "1"},
	}
	for _, test := range tests {
		if got := addBinary(test.num1, test.num2); got != test.want {
			t.Errorf("addBinary(%s, %s) = %s, want %s", test.num1, test.num2, got, test.want)
		}
	}
}
