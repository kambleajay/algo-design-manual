package revpolinot

import "strconv"

type Stack struct {
	items []int
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func eval(n1, n2 int, operator string) int {
	var result int
	switch operator {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		result = n1 / n2
	}
	return result
}

func evalRPN(tokens []string) int {
	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	nums := Stack{}
	for _, token := range tokens {
		if !operators[token] {
			n, _ := strconv.Atoi(token)
			nums.push(n)
		} else {
			n1 := nums.pop()
			n2 := nums.pop()
			nums.push(eval(n2, n1, token))
		}
	}
	return nums.pop()
}
