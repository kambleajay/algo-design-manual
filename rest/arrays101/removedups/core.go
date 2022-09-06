package removedups

func removeDuplicates(nums []int) int {
	var reader, writer int
	N := len(nums)
	for ; reader < N; reader++ {
		if reader == 0 || nums[reader] != nums[reader-1] {
			nums[writer] = nums[reader]
			writer++
		}
	}
	return writer
}
