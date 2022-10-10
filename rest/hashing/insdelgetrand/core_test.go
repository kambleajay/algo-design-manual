package insdelgetrand

import "testing"

type Insert struct {
	num  int
	want bool
}

type Remove struct {
	num  int
	want bool
}

type GetRandom struct {
	set []int
}

func notIn(x int, ys []int) bool {
	for _, y := range ys {
		if x == y {

			return false
		}
	}
	return true
}

func TestRandomizedSet(t *testing.T) {
	var tests = [][]interface{}{
		{Insert{1, true}, Remove{2, false}, Insert{2, true},
			GetRandom{[]int{1, 2}}, Remove{1, true}, Insert{2, false},
			GetRandom{[]int{2}}},
		{Insert{0, true}, Insert{1, true}, Remove{0, true},
			Insert{2, true}, Remove{1, true}, GetRandom{[]int{2}}},
	}
	for i, test := range tests[1:] {
		rs := Constructor()
		for _, command := range test {
			switch command := command.(type) {
			case Insert:
				if got := rs.Insert(command.num); got != command.want {
					t.Errorf("Test %d: Insert(%d) = %t, want %t", i, command.num, got, command.want)
				}
			case Remove:
				if got := rs.Remove(command.num); got != command.want {
					t.Errorf("Test %d: Remove(%d) = %t, want %t", i, command.num, got, command.want)
				}
			case GetRandom:
				if got := rs.GetRandom(); notIn(got, command.set) {
					t.Errorf("Test %d: GetRandom() = %d, not in %v", i, got, command.set)
				}
			}
		}
	}
}
