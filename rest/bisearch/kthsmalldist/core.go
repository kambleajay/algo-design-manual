package kthsmalldist

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	N := len(nums)
	loSum, hiSum := 0, nums[N-1]-nums[0]
	for loSum < hiSum {
		midSum := loSum + (hiSum-loSum)/2
		var count, left int
		for right := 0; right < N; right++ {
			for nums[right]-nums[left] > midSum {
				left++
			}
			count += right - left
		}
		if count >= k {
			hiSum = midSum
		} else {
			loSum = midSum + 1
		}
	}
	return loSum
}
