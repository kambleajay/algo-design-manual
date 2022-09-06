package movezeros

func moveZeroes(nums []int) {
	N := len(nums)
	if N == 0 {
		return
	}
	var reader, writer int
	for ; reader < N; reader++ {
		if nums[reader] != 0 {
			nums[writer], nums[reader] = nums[reader], nums[writer]
			writer++
		}
	}
}
