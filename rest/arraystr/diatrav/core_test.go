package diatrav

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestDiagonalTraversal(t *testing.T) {
	var tests = []struct {
		input [][]int
		want  []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		{[][]int{{1}}, []int{1}},
		{[][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 3, 5, 4, 6}},
	}
	for _, test := range tests {
		if got := findDiagonalOrder(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("findDiagonalOrder(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func BenchmarkDiagonalTraversal(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	m := 10000
	n := 10000
	input := make([][]int, m)
	min, max := int(-1e5), int(1e5)
	for i := 0; i < m; i++ {
		input[i] = make([]int, n)
		for j := 0; j < n; j++ {
			input[i][j] = rng.Intn(max-min) + min
		}
	}
	for i := 0; i < b.N; i++ {
		findDiagonalOrder(input)
	}
}
