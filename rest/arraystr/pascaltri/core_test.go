package pascaltri

import (
	"reflect"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	var tests = []struct {
		rows int
		want [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{1, [][]int{{1}}},
	}
	for _, test := range tests {
		if got := generate(test.rows); !reflect.DeepEqual(got, test.want) {
			t.Errorf("generate(%d) = %v, want %v", test.rows, got, test.want)
		}
	}
}
