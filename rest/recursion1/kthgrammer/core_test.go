package kthgrammer

import "testing"

func TestKthGrammer(t *testing.T) {
	var tests = []struct {
		n, k, want int
	}{
		{1, 1, 0},
		{2, 1, 0},
		{2, 2, 1},
		{4, 5, 1},
		{5, 9, 1},
		{5, 6, 0},
		{5, 15, 1},
		{5, 16, 0},
		{30, 434991989, 0},
	}
	for _, test := range tests {
		if got := kthGrammar(test.n, test.k); got != test.want {
			t.Errorf("kthGrammar(%d, %d) = %d, want %d", test.n, test.k, got, test.want)
		}
	}
}

func BenchmarkKthGrammer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kthGrammar(30, 434991989)
	}
}
