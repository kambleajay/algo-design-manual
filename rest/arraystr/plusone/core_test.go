package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}
	for _, test := range tests {
		if got := plusOne(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("plusOne(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
