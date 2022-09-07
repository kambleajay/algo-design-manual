package maxconones

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxConsecutiveOnes(nums []int) int {
	N := len(nums)
	var longestSeq, left, right, numZeros int
	for ; right < N; right++ {
		if nums[right] == 0 {
			numZeros++
		}
		for numZeros == 2 {
			if nums[left] == 0 {
				numZeros--
			}
			left++
		}
		longestSeq = max(longestSeq, right-left+1)
	}
	return longestSeq
}
