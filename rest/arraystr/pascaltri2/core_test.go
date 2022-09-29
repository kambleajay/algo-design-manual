package pascaltri2

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	var tests = []struct {
		rowIndex int
		want     []int
	}{
		{0, []int{1}},
		{1, []int{1, 1}},
		{2, []int{1, 2, 1}},
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
	}
	for _, test := range tests {
		if got := getRow(test.rowIndex); !reflect.DeepEqual(got, test.want) {
			t.Errorf("getRow(%d) = %v, want %v", test.rowIndex, got, test.want)
		}
	}
}
