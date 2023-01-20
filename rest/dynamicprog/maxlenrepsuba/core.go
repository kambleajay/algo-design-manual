package maxlenrepsuba

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	ans := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans
}
