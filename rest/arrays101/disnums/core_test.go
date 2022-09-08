package disnums

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		{[]int{3, 3, 2, 1, 4, 5, 6, 4}, []int{7, 8}},
		{[]int{1, 1}, []int{2}},
		{[]int{}, []int{}},
	}
	for _, test := range tests {
		if got := findDisappearedNumbers(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("findDisappearedNumbers(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
