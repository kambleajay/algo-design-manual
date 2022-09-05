package runningsum

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
		{[]int{}, []int{}},
		{nil, nil},
	}
	for _, test := range tests {
		if got := runningSum(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("runningSum(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
