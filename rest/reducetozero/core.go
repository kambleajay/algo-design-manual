package reducetozero

func numberOfSteps(num int) int {
	var steps int
	for num != 0 {
		if (num & 1) == 0 {
			num = num >> 1
		} else {
			num--
		}
		steps++
	}
	return steps
}
