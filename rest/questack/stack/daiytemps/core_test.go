package daiytemps

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	var tests = []struct {
		temps []int
		want  []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}
	for _, test := range tests {
		if got := dailyTemperatures(test.temps); !reflect.DeepEqual(got, test.want) {
			t.Errorf("dailyTemperatures(%v) = %v, want %v", test.temps, got, test.want)
		}
	}
}
