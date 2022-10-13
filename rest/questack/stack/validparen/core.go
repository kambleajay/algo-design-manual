package validparen

import "fmt"

type Stack struct {
	items []rune
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) push(r rune) {
	s.items = append(s.items, r)
}

func (s *Stack) pop() rune {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func isValid(s string) bool {
	open := Stack{}
	openBrackets := map[rune]bool{'{': true, '[': true, '(': true}
	validSet := map[string]bool{"{}": true, "[]": true, "()": true}
	for _, r := range s {
		if openBrackets[r] {
			open.push(r)
		} else {
			if open.isEmpty() {
				return false
			}
			lastOpen := open.pop()
			together := fmt.Sprintf("%c%c", lastOpen, r)
			if validSet[together] {
				continue
			} else {
				return false
			}
		}
	}
	return open.isEmpty()
}
