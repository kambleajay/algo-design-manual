package minindsum

import (
	"reflect"
	"testing"
)

func TestFindRestaurant(t *testing.T) {
	var tests = []struct {
		list1 []string
		list2 []string
		want  []string
	}{
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}, []string{"Shogun"}},
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}, []string{"Shogun"}},
		{[]string{"happy", "sad", "good"}, []string{"sad", "happy", "good"}, []string{"sad", "happy"}},
	}
	for _, test := range tests {
		if got := findRestaurant(test.list1, test.list2); !reflect.DeepEqual(got, test.want) {
			t.Errorf("findRestaurant(%v, %v) = %v, want %v", test.list1, test.list2, got, test.want)
		}
	}
}
