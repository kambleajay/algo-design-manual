package rotate

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}
