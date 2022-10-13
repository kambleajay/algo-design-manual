package daiytemps

type Temperature struct {
	day  int
	temp int
}

type Stack struct {
	items []Temperature
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) push(t Temperature) {
	s.items = append(s.items, t)
}

func (s *Stack) peek() Temperature {
	if s.isEmpty() {
		panic("peek called on an empty stack")
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) pop() Temperature {
	if s.isEmpty() {
		panic("pop called on an empty stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func dailyTemperatures(temperatures []int) []int {
	stack := Stack{}
	answer := make([]int, len(temperatures))
	for day, temp := range temperatures {
		for !stack.isEmpty() && stack.peek().temp < temp {
			last := stack.peek()
			answer[last.day] = day - last.day
			stack.pop()
		}
		stack.push(Temperature{day, temp})
	}
	return answer
}
