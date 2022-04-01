package remove_k_digits

import "testing"

func TestRemoveKDigits(t *testing.T) {
	var tests = []struct {
		num  string
		k    int
		want string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
		{"", 0, "0"},
		{"123", 3, "0"},
		{"366431814", 4, "31814"},
		{"595103317", 4, "3317"},
		{"11111", 3, "11"},
		{"11122", 2, "111"},
		{"54618810", 2, "418810"},
		{"408461689", 7, "1"},
		{"518870579", 6, "57"},
		{"150628542", 2, "628542"},
		{"10", 1, "0"},
		{"10001", 4, "0"},
	}
	for _, test := range tests {
		if got := removeKdigits(test.num, test.k); got != test.want {
			t.Errorf("removeKdigits(%s, %d) should be %s, but was %s", test.num, test.k, test.want, got)
		}
	}
}
