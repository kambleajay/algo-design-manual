package disnums

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	for _, n := range nums {
		seenIndex := abs(n) - 1
		if nums[seenIndex] > 0 {
			nums[seenIndex] = -nums[seenIndex]
		}
	}
	var answer []int
	for i, n := range nums {
		if n > 0 {
			answer = append(answer, i+1)
		}
	}
	return answer
}
