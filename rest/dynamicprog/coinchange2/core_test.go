package coinchange2

import "testing"

func TestChange(t *testing.T) {
	var tests = []struct {
		amount int
		coins  []int
		want   int
	}{
		{5, []int{1, 2, 5}, 4},
		{3, []int{2}, 0},
		{10, []int{10}, 1},
	}
	for _, test := range tests {
		if got := change(test.amount, test.coins); got != test.want {
			t.Errorf("change(%d, %v) = %d, want %d", test.amount, test.coins, got, test.want)
		}
	}
}

func BenchmarkChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		change(11, []int{2, 5, 10})
	}
}
