package searchsizeunk

import "testing"

func TestSearch(t *testing.T) {
	var tests = []struct {
		items  []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}
	for _, test := range tests {
		ar := NewArrayReader(test.items)
		if got := search(ar, test.target); got != test.want {
			t.Errorf("search(ArrayReader(%v), %d) = %d, want %d", test.items, test.target, got, test.want)
		}
	}
}
