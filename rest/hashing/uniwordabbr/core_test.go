package uniwordabbr

import "testing"

type Want struct {
	word string
	want bool
}

func TestIsUnique(t *testing.T) {
	var tests = []struct {
		words []string
		wants []Want
	}{
		{[]string{"deer", "door", "cake", "card"},
			[]Want{{"dear", false}, {"cart", true}, {"cane", false}, {"make", true}, {"cake", true}}},
		{[]string{"deer", "door", "cake", "card"},
			[]Want{{"dear", false}, {"door", false}, {"cart", true}, {"cake", true}}},
	}
	for i, test := range tests {
		dict := Constructor(test.words)
		for _, want := range test.wants {
			if got := dict.IsUnique(want.word); got != want.want {
				t.Errorf("[%d] IsUnique(%s) = %t, want %t", i, want.word, got, want.want)
			}
		}
	}

}
