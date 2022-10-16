package quewithstacks

import "testing"

type Push struct {
	num int
}

type Pop struct {
	want int
}

type Peek struct {
	want int
}

type Empty struct {
	want bool
}

func TestQueueWithStacks(t *testing.T) {
	var tests = [][]interface{}{
		{Push{1}, Push{2}, Peek{1}, Pop{1}, Empty{false}, Peek{2}, Pop{2},
			Empty{true}, Push{3}, Push{4}, Push{5}, Empty{false}, Pop{3},
			Pop{4}, Push{6}, Pop{5}, Empty{false}, Pop{6}, Empty{true}},
	}
	for _, test := range tests {
		q := Constructor()
		for _, inst := range test {
			switch inst := inst.(type) {
			case Push:
				q.Push(inst.num)
			case Peek:
				if got := q.Peek(); got != inst.want {
					t.Errorf("Peek() = %d, want %d", got, inst.want)
				}
			case Pop:
				if got := q.Pop(); got != inst.want {
					t.Errorf("Pop() = %d, want %d", got, inst.want)
				}
			case Empty:
				if got := q.Empty(); got != inst.want {
					t.Errorf("Empty() = %t, want %t", got, inst.want)
				}
			}
		}
	}
}
