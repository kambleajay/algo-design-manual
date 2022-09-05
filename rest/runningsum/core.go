package runningsum

func runningSum(nums []int) []int {
	N := len(nums)
	if N == 0 || N == 1 {
		return nums
	}
	for i := 1; i < N; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
