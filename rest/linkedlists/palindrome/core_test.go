package palindrome

import (
	"algo/rest/linkedlists/utils"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input []int
		want  bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3}, false},
		{[]int{1, 2, 1}, true},
		{[]int{}, true},
	}
	for _, test := range tests {
		list := utils.ToList(test.input...)
		if got := isPalindrome(list); got != test.want {
			t.Errorf("isPalindrome(%v) = %t, want %t", list, got, test.want)
		}
	}
}
