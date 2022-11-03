package kthsmallelesortmat

import "testing"

func TestKthSmallest(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13},
		{[][]int{{-5}}, 1, -5},
		{[][]int{{1, 2}, {1, 3}}, 2, 1},
	}
	for _, test := range tests {
		if got := kthSmallest(test.matrix, test.k); got != test.want {
			t.Errorf("kthSmallest(%v, %d) = %d, want %d", test.matrix, test.k, got, test.want)
		}
	}
}
