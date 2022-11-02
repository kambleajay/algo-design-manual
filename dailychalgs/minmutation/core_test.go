package minmutation

import "testing"

func TestMinMutation(t *testing.T) {
	var tests = []struct {
		start string
		end   string
		bank  []string
		want  int
	}{
		{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},
		{"AACCGGTT", "AACCGGTA", []string{}, -1},
		{"AACCGGTT", "AAACGGTA", []string{"AACCGATT", "AACCGATA", "AAACGATA", "AAACGGTA"}, 4},
		{"AAAAAAAA", "CCCCCCCC", []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA"}, -1},
	}
	for _, test := range tests {
		if got := minMutation(test.start, test.end, test.bank); got != test.want {
			t.Errorf("minMutation(%s, %s, %v) = %d, want %d", test.start, test.end, test.bank, got, test.want)
		}
	}
}
