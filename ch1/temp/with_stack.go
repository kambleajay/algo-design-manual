package temp

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	var stack []int
	for day, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			prevDay := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[prevDay] = day - prevDay
		}
		stack = append(stack, day)
	}
	return answer
}
