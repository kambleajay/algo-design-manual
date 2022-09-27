package pivotsum

func pivotIndex(nums []int) int {
	var sum, leftSum, rightSum int
	for _, n := range nums {
		sum += n
	}
	for i, n := range nums {
		rightSum = sum - leftSum - n
		if leftSum == rightSum {
			return i
		}
		leftSum += n
	}
	return -1
}
