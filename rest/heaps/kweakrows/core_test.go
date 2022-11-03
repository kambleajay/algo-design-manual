package kweakrows

import (
	"reflect"
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	var tests = []struct {
		mat  [][]int
		k    int
		want []int
	}{
		{[][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3, []int{2, 0, 3}},
		{[][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 2, []int{0, 2}},
		{[][]int{{0}, {0}, {0}}, 2, []int{0, 1}},
	}
	for _, test := range tests {
		if got := kWeakestRows(test.mat, test.k); !reflect.DeepEqual(got, test.want) {
			t.Errorf("kWeakestRows(%v, %d) = %v, want %v", test.mat, test.k, got, test.want)
		}
	}
}
