package nextgreatlet

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	var tests = []struct {
		letters      []byte
		target, want byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'x', 'x', 'y', 'y'}, 'z', 'x'},
		{[]byte{'x', 'x', 'y', 'y'}, 'y', 'x'},
		{[]byte{'x', 'x', 'y', 'y'}, 'x', 'y'},
	}
	for _, test := range tests {
		if got := nextGreatestLetter(test.letters, test.target); got != test.want {
			t.Errorf("nextGreatestLetter(%v, %c) = %c, want %c", test.letters, test.target, got, test.want)
		}
	}
}
