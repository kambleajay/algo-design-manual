package stackwithques

import "testing"

type Push struct {
	x int
}

type Top struct {
	want int
}

type Pop struct {
	want int
}

type Empty struct {
	want bool
}

func TestStackWithQueues(t *testing.T) {
	var tests = [][]interface{}{
		{Push{1}, Push{2},
			Top{2}, Pop{2}, Empty{false},
			Push{3},
			Top{3}, Pop{3}, Empty{false},
			Top{1}, Pop{1}, Empty{true},
			Push{4}, Empty{false},
			Top{4}, Pop{4}, Empty{true}},
	}
	for _, test := range tests {
		stack := Constructor()
		for _, command := range test {
			switch command := command.(type) {
			case Push:
				stack.Push(command.x)
			case Top:
				if got := stack.Top(); got != command.want {
					t.Errorf("Top() = %d, want %d", got, command.want)
				}
			case Pop:
				if got := stack.Pop(); got != command.want {
					t.Errorf("Pop() = %d, want %d", got, command.want)
				}
			case Empty:
				if got := stack.Empty(); got != command.want {
					t.Errorf("Empty() = %t, want %t", got, command.want)
				}
			}
		}
	}
}
