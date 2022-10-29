package zeroonemat

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	var tests = []struct {
		input [][]int
		want  [][]int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
		{[][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}, {0, 1, 0}}},
	}
	for _, test := range tests {
		if got := updateMatrix(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("updateMatrix(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
