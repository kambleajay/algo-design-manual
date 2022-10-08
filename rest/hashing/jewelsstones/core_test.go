package jewelsstones

import "testing"

func TestNumOfJewelsInStones(t *testing.T) {
	var tests = []struct {
		jewels string
		stones string
		want   int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, test := range tests {
		if got := numJewelsInStones(test.jewels, test.stones); got != test.want {
			t.Errorf("numJewelsInStones(%s, %s) = %d, want %d", test.jewels, test.stones, got, test.want)
		}
	}
}
