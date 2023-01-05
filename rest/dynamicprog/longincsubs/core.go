package longincsubs

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLISDP(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	answer := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
				answer = max(answer, dp[i])
			}
		}
	}
	return answer
}

func binarySearch(sub []int, num int) int {
	lo, hi := 0, len(sub)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if num == sub[mid] {
			return mid
		} else if num < sub[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func lengthOfLISBinarySearch(nums []int) int {
	sub := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > sub[len(sub)-1] {
			sub = append(sub, nums[i])
		} else {
			j := binarySearch(sub, nums[i])
			sub[j] = nums[i]
		}
	}
	return len(sub)
}

func lengthOfLIS(nums []int) int {
	return lengthOfLISBinarySearch(nums)
}
