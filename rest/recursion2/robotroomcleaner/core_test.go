package robotroomcleaner

import "testing"

func TestCleanRoom(t *testing.T) {
	var tests = []struct {
		rooms    [][]int
		row, col int
		want     int
	}{
		{[][]int{{1}}, 0, 0, 1},
		{[][]int{{1, 0}, {1, 1}}, 1, 1, 3},
		{[][]int{{1, 1, 1, 1, 1, 0, 1, 1}, {1, 1, 1, 1, 1, 0, 1, 1}, {1, 0, 1, 1, 1, 1, 1, 1}, {0, 0, 0, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}}, 1, 3, 30},
	}
	for _, test := range tests {
		robot := NewRobot(test.rooms, test.row, test.col)
		cleanRoom(&robot)
		if len(robot.cleaned) != test.want {
			t.Errorf("cleanRoom() cleaned %d rooms, want %d", len(robot.cleaned), test.want)
		}
	}
}
