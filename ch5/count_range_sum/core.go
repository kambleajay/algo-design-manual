package count_range_sum

func countWithMergeSort(sums []int, lo, hi, lower, upper int) int {
	if hi-lo <= 1 {
		return 0
	}
	mid := (lo + hi) / 2
	count := countWithMergeSort(sums, lo, mid, lower, upper) + countWithMergeSort(sums, mid, hi, lower, upper)
	j, k, t := mid, mid, mid
	aux := make([]int, hi-lo)
	for i, r := lo, 0; i < mid; i, r = i+1, r+1 {
		for k < hi && sums[k]-sums[i] < lower {
			k++
		}
		for j < hi && sums[j]-sums[i] <= upper {
			j++
		}
		for t < hi && sums[t] < sums[i] {
			aux[r] = sums[t]
			r++
			t++
		}
		aux[r] = sums[i]
		count += j - k
	}
	copy(sums[lo:t], aux)
	return count
}

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	return countWithMergeSort(sums, 0, n+1, lower, upper)
}
