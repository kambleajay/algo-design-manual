package arrpart1

import "sort"

func arrayPairSum1(nums []int) int {
	K := int(1e4)
	size := (2 * K) + 1
	freqs := make([]int, size)
	for _, num := range nums {
		freqs[num+K]++
	}
	isEvenIndex := true
	var maxSum int
	for num, freq := range freqs {
		for freq > 0 {
			if isEvenIndex {
				maxSum += num - K
			}
			isEvenIndex = !isEvenIndex
			freq--
		}
	}
	return maxSum
}

func arrayPairSum2(nums []int) int {
	sort.Ints(nums)
	maxSum := 0
	for i := 0; i < len(nums); i += 2 {
		maxSum += nums[i]
	}
	return maxSum
}

func arrayPairSum(nums []int) int {
	return arrayPairSum1(nums)
}
