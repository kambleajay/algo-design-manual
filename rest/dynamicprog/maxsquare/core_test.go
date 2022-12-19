package maxsquare

import "testing"

func TestMaximalSquare(t *testing.T) {
	var tests = []struct {
		matrix [][]byte
		want   int
	}{
		{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, 4},
		{[][]byte{{'0', '1'}, {'1', '0'}}, 1},
		{[][]byte{{'0'}}, 0},
	}
	for _, test := range tests {
		if got := maximalSquare(test.matrix); got != test.want {
			t.Errorf("maximalSquare(%v) = %d, want %d", test.matrix, got, test.want)
		}
	}
}
