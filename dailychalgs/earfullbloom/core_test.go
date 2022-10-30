package earfullbloom

import "testing"

func TestEarliestFullBloom(t *testing.T) {
	var tests = []struct {
		plantTimes []int
		growTimes  []int
		want       int
	}{
		{[]int{1, 4, 3}, []int{2, 3, 1}, 9},
		{[]int{1, 2, 3, 2}, []int{2, 1, 2, 1}, 9},
		{[]int{1}, []int{1}, 2},
		{[]int{3, 11, 29, 4, 4, 26, 26, 12, 13, 10, 30, 19, 27, 2, 10}, []int{10, 13, 22, 17, 18, 15, 21, 11, 24, 14, 18, 23, 1, 30, 6}, 227},
	}
	for _, test := range tests {
		if got := earliestFullBloom(test.plantTimes, test.growTimes); got != test.want {
			t.Errorf("earliestFullBloom(%v, %v) = %d, want %d", test.plantTimes, test.growTimes, got, test.want)
		}
	}
}
