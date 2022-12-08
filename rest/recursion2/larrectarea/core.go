package larrectarea

type Stack struct {
	items []int
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) peek() int {
	return s.items[len(s.items)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func largestRectangleArea(heights []int) int {
	var stack Stack
	stack.push(-1)
	n := len(heights)
	var maxArea int
	for k := 0; k < n; k++ {
		for stack.peek() != -1 && heights[stack.peek()] >= heights[k] {
			height := heights[stack.pop()]
			width := k - stack.peek() - 1
			maxArea = max(maxArea, width*height)
		}
		stack.push(k)
	}
	for stack.peek() != -1 {
		height := heights[stack.pop()]
		width := n - stack.peek() - 1
		maxArea = max(maxArea, width*height)
	}
	return maxArea
}
