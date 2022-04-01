package two_sum_2

func twoSum(numbers []int, target int) []int {
	for lo, hi := 0, len(numbers)-1; lo <= hi; {
		sum := numbers[lo] + numbers[hi]
		if sum > target {
			hi--
		} else if sum < target {
			lo++
		} else {
			return []int{lo + 1, hi + 1}
		}
	}
	return nil
}
