package remelems

import (
	"algo/rest/linkedlists/utils"
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	var tests = []struct {
		list         []int
		elemToRemove int
		want         []int
	}{
		{[]int{1, 2, 6, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
		{[]int{}, 1, nil},
		{[]int{7, 7, 7, 7}, 7, nil},
	}
	for _, test := range tests {
		if got := removeElements(utils.ToList(test.list...), test.elemToRemove); !reflect.DeepEqual(utils.ToSlice(got), test.want) {
			t.Errorf("removeElements(%v, %d) = %v, want %v", test.list, test.elemToRemove, got, test.want)
		}
	}
}
