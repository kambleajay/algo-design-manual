package temp

func dailyTemperatures3(temperatures []int) []int {
	N := len(temperatures)
	answer := make([]int, N)
	hottest := 0
	for i := N - 1; i >= 0; i-- {
		if temperatures[i] >= hottest {
			hottest = temperatures[i]
			continue
		}
		daysToWarmer := 1
		for temperatures[i+daysToWarmer] <= temperatures[i] {
			daysToWarmer += answer[i+daysToWarmer]
		}
		answer[i] = daysToWarmer
	}
	return answer
}
