package firunich

import "testing"

func TestFirstUniqueChar(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}
	for _, test := range tests {
		if got := firstUniqChar(test.s); got != test.want {
			t.Errorf("firstUniqChar(%s) = %d, want %d", test.s, got, test.want)
		}
	}
}
