package reconstruct_queue

import (
	"reflect"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	var tests = []struct {
		input [][]int
		want  [][]int
	}{
		{[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}, [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}},
		{[][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}, [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}}},
		{[][]int{{5, 0}}, [][]int{{5, 0}}},
	}
	for _, test := range tests {
		copyOfInput := make([][]int, len(test.input))
		copy(copyOfInput, test.input)
		if got := reconstructQueue(test.input); !reflect.DeepEqual(test.want, got) {
			t.Errorf("reconstructQueue(%v) should be %v, but was %v", copyOfInput, test.want, got)
		}
	}
}
