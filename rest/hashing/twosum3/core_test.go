package twosum3

import "testing"

type Instruction struct {
	opType string
	num    int
	want   bool
}

func add(num int) Instruction {
	return Instruction{opType: "add", num: num}
}

func find(num int, want bool) Instruction {
	return Instruction{"find", num, want}
}

func TestTwoSum(t *testing.T) {
	var tests = [][]Instruction{
		{add(1), add(3), add(5), find(4, true), find(7, false), add(6), find(7, true)},
		{add(0), find(0, false)},
		{add(1), find(2, false), add(1), find(2, true)},
	}
	for _, test := range tests {
		twoSum := Constructor()
		for _, ins := range test {
			switch ins.opType {
			case "add":
				twoSum.Add(ins.num)
			case "find":
				if got := twoSum.Find(ins.num); got != ins.want {
					t.Errorf("twoSum.Find(%d) = %t, want %t", ins.num, got, ins.want)
				}
			}
		}
	}
}
