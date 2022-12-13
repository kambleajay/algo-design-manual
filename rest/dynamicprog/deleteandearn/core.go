package deleteandearn

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func deleteAndEarn(nums []int) int {
	points := make(map[int]int)
	var maxNum int
	for _, num := range nums {
		points[num] += num
		maxNum = max(maxNum, num)
	}
	dp := make([]int, maxNum+1)
	dp[1] = points[1]
	for k := 2; k <= maxNum; k++ {
		take := points[k] + dp[k-2]
		notTake := dp[k-1]
		dp[k] = max(take, notTake)
	}
	return dp[maxNum]
}
