package movavg

import "testing"

type MovingAverageInstruction struct {
	size int
}

type NextInstruction struct {
	val  int
	want float64
}

func TestMovingAverage(t *testing.T) {
	var tests = [][]interface{}{
		{MovingAverageInstruction{3}, NextInstruction{1, 1.0}, NextInstruction{10, 5.5},
			NextInstruction{3, 4.66667}, NextInstruction{5, 6.0}},
	}
	for _, test := range tests {
		var ma MovingAverage
		for _, inst := range test {
			switch inst := inst.(type) {
			case MovingAverageInstruction:
				ma = Constructor(inst.size)
			case NextInstruction:
				if got := ma.Next(inst.val); got != inst.want {
					t.Errorf("Next(%d) = %f, want %f", inst.val, got, inst.want)
				}
			}
		}
	}
}
