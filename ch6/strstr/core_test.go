package strstr

import "testing"

func TestStrStr(t *testing.T) {
	var tests = []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"abc", "", 0},
		{"", "", 0},
		{"findinahaystackneedleina", "needle", 15},
		{"abcde", "a", 0},
		{"aaaa", "aaa", 0},
		{"aaa", "aaaa", -1},
		{"aaa", "aaa", 0},
		{"abaaabbabbaabaababaababbabababaababbbababbaababbabbbbaababaaaaabababbbaaaabbbaabaaababbaabaaabaaaaaaaaaaababaaaaabbbaabaaaaabaabbabbaabbbbababbaabbabbabbabababaabbbbaaaaaaabbabaababbbaab", "bbbbbaaabaaaabaaabbaabbaabbabaaabbbbabbbb", -1},
	}
	for _, test := range tests {
		if got := strStr(test.haystack, test.needle); got != test.want {
			t.Errorf("strStr(%s, %s) was %d, it should have been %d", test.haystack, test.needle, got, test.want)
		}
	}
}

func BenchmarkStrStr(b *testing.B) {
	haystack := "abaaabbabbaabaababaababbabababaababbbababbaababbabbbbaababaaaaabababbbaaaabbbaabaaababbaabaaabaaaaaaaaaaababaaaaabbbaabaaaaabaabbabbaabbbbababbaabbabbabbabababaabbbbaaaaaaabbabaababbbaab"
	needle := "bbbbbaaabaaaabaaabbaabbaabbabaaabbbbabbbb"
	for i := 0; i < b.N; i++ {
		strStr(haystack, needle)
	}
}
