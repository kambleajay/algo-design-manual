package heapsort

func less(a []int, i, j int) bool {
	return a[i-1] < a[j-1]
}

func greater(a []int, i, j int) bool {
	return a[i-1] > a[j-1]
}

func exch(a []int, i, j int) {
	a[i-1], a[j-1] = a[j-1], a[i-1]
}

func sink(a []int, k, N int) {
	for 2*k <= N {
		j := 2 * k
		if j < N && less(a, j, j+1) {
			j++
		}
		if greater(a, k, j) {
			break
		}
		exch(a, k, j)
		k = j
	}
}

func heapSort(nums []int) []int {
	N := len(nums)
	for k := N / 2; k > 1; k-- {
		sink(nums, k, N)
	}
	for N > 1 {
		exch(nums, 1, N)
		N--
		sink(nums, 1, N)
	}
	return nums
}
