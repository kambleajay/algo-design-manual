package wallsgates

import (
	"reflect"
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	max := 2147483647
	var tests = []struct {
		rooms [][]int
		want  [][]int
	}{
		{[][]int{{max, -1, 0, max}, {max, max, max, -1}, {max, -1, max, -1}, {0, -1, max, max}},
			[][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}}},
	}
	for _, test := range tests {
		copyOfRooms := make([][]int, len(test.rooms))
		for i, row := range test.rooms {
			copyOfRooms[i] = make([]int, len(row))
			copy(copyOfRooms[i], row)
		}
		if wallsAndGates(test.rooms); !reflect.DeepEqual(test.rooms, test.want) {
			t.Errorf("wallsAndGates(%v) = %v, want %v", copyOfRooms, test.rooms, test.want)
		}
	}
}
