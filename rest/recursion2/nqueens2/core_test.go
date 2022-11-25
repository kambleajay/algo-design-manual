package nqueens2

import "testing"

func TestTotalNQueens(t *testing.T) {
	var tests = []struct {
		n, want int
	}{
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
	}
	for _, test := range tests {
		if got := totalNQueens(test.n); got != test.want {
			t.Errorf("totalNQueens(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
