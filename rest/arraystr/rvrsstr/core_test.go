package rvrsstr

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	var tests = []struct {
		s    []byte
		want []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Hannah"), []byte("hannaH")},
	}
	for _, test := range tests {
		copyOfS := make([]byte, len(test.s))
		if reverseString(test.s); !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("reverseString(%v) = %v, want %v", copyOfS, test.s, test.want)
		}
	}
}
