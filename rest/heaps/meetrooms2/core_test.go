package meetrooms2

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	var tests = []struct {
		interval [][]int
		want     int
	}{
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{[][]int{{7, 10}, {2, 4}}, 1},
		{[][]int{{10, 20}}, 1},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 1},
		{[][]int{{1, 3}, {2, 4}, {3, 6}, {5, 8}}, 2},
	}
	for _, test := range tests {
		if got := minMeetingRooms(test.interval); got != test.want {
			t.Errorf("minMeetingRooms(%v) = %d, want %d", test.interval, got, test.want)
		}
	}
}
