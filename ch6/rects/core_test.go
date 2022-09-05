package rects

import (
	"fmt"
	"testing"
)

func inRect(point []int, rect []int) bool {
	px, py := point[0], point[1]
	x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
	return (px >= x1) && (px <= x2) && (py >= y1) && (py <= y2)
}

func TestSolutionPick(t *testing.T) {
	var tests = []struct {
		rects [][]int
	}{
		{[][]int{{-2, -2, 1, 1}, {2, 2, 4, 6}}},
		{[][]int{{82918473, -57180867, 82918476, -57180863}, {83793579, 18088559, 83793580, 18088560}, {66574245, 26243152, 66574246, 26243153}, {72983930, 11921716, 72983934, 11921720}}},
	}
	for _, test := range tests {
		solution := Constructor(test.rects)
		rectFreq := make(map[int]int)
		for i := 0; i < 1000000; i++ {
			point := solution.Pick()
			x, y := point[0], point[1]
			found := false
			for i, r := range test.rects {
				inR := inRect(point, r)
				if inR {
					rectFreq[i]++
					found = found || inR
				}
			}
			if !found {
				t.Errorf("point (%d, %d) is not in any rectangle", x, y)
			}
		}
		var totalCount int
		for _, c := range rectFreq {
			totalCount += c
		}
		for r, c := range rectFreq {
			fmt.Printf("%d = %.2f%%, ", r, (float64(c)/float64(totalCount))*100)
		}
		fmt.Println("")
	}
}
