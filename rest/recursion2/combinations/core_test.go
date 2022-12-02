package combinations

import (
	"algo/testing/utils"
	"testing"
)

func TestCombinations(t *testing.T) {
	var tests = []struct {
		n, k int
		want [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{1, 1, [][]int{{1}}},
		{5, 3, [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}}},
	}
	for _, test := range tests {
		if got := combine(test.n, test.k); !utils.ContainsAll(got, test.want) {
			t.Errorf("combine(%d, %d) = %v, want %v", test.n, test.k, got, test.want)
		}
	}
}
