package remove_k_digits

import "strings"

type Stack struct {
	a []uint8
}

func (s *Stack) push(v uint8) {
	s.a = append(s.a, v)
}

func (s *Stack) top() uint8 {
	return s.a[len(s.a)-1]
}

func (s *Stack) pop() uint8 {
	top := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return top
}

func (s *Stack) isEmpty() bool {
	return s.a == nil || len(s.a) == 0
}

func removeKdigits(num string, k int) string {
	var stack Stack
	for i := 0; i < len(num); i++ {
		for !stack.isEmpty() && k > 0 && stack.top() > num[i] {
			stack.pop()
			k--
		}
		stack.push(num[i])
	}
	for k > 0 {
		stack.pop()
		k--
	}
	answer := string(stack.a)
	answer = strings.TrimLeft(answer, "0")
	if len(answer) == 0 {
		return "0"
	}
	return answer
}
