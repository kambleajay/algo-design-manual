package reversestr

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"hello", "olleh"},
		{"Hannah", "hannaH"},
	}
	for _, test := range tests {
		ba := []byte(test.s)
		copyOfBA := make([]byte, len(ba))
		copy(copyOfBA, ba)
		if reverseString(ba); !reflect.DeepEqual(ba, []byte(test.want)) {
			t.Errorf("reverseString(%v) = %v, want %v", copyOfBA, test.s, test.want)
		}
	}
}
