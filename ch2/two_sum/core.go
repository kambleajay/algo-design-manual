package two_sum

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)
	for i, n := range nums {
		x := target - n
		indexOfX, ok := numToIndex[x]
		if ok {
			return []int{indexOfX, i}
		}
		numToIndex[n] = i
	}
	return nil
}
