package hashmap

import (
	"reflect"
	"testing"
)

type Test struct {
	operation string
	operands  []int
	want      interface{}
}

func set1() []Test {
	return []Test{
		Test{"put", []int{1, 1}, nil},
		Test{"put", []int{2, 2}, nil},
		Test{"get", []int{1}, 1},
		Test{"get", []int{3}, -1},
		Test{"put", []int{2, 1}, nil},
		Test{"get", []int{2}, 1},
		Test{"remove", []int{2}, nil},
		Test{"get", []int{2}, -1},
	}
}

func TestHashMap(t *testing.T) {
	var testSets = [][]Test{set1()}

	for _, set := range testSets {
		m := Constructor()
		for _, test := range set {
			switch test.operation {
			case "put":
				k, v := test.operands[0], test.operands[1]
				m.Put(k, v)
			case "get":
				got := m.Get(test.operands[0])
				if !reflect.DeepEqual(got, test.want) {
					t.Errorf("m.Get(%d) = %v, want %v", test.operands[0], got, test.want)
				}
			case "remove":
				m.Remove(test.operands[0])
			}
		}
	}
}
