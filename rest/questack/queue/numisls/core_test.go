package numisls

import "testing"

func TestNumIslands(t *testing.T) {
	var tests = []struct {
		grid [][]byte
		want int
	}{
		{[][]byte{{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'}}, 1},
		{[][]byte{{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '0'}}, 3},
	}
	for i, test := range tests {
		if got := numIslands(test.grid); got != test.want {
			t.Errorf("[%d] numIslands(%v) = %d, want %d", i, test.grid, got, test.want)
		}
	}
}
