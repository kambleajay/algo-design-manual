package fldfill

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	var tests = []struct {
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0, [][]int{{0, 0, 0}, {0, 0, 0}}},
	}
	for _, test := range tests {
		if got := floodFill(test.image, test.sr, test.sc, test.color); !reflect.DeepEqual(got, test.want) {
			t.Errorf("floodFill(%v, %d, %d, %d) = %v, want %v", test.image, test.sr, test.sc, test.color, got, test.want)
		}
	}
}
