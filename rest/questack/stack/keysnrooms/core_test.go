package keysnrooms

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	var tests = []struct {
		rooms [][]int
		want  bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
	}
	for _, test := range tests {
		if got := canVisitAllRooms(test.rooms); got != test.want {
			t.Errorf("canVisitAllRooms(%v) = %t, want %t", test.rooms, got, test.want)
		}
	}
}
