package sortarr

func merge(a, aux []int, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}
	left, right := lo, mid+1
	for writer := lo; writer <= hi; writer++ {
		if left > mid {
			a[writer] = aux[right]
			right++
		} else if right > hi {
			a[writer] = aux[left]
			left++
		} else if aux[right] < aux[left] {
			a[writer] = aux[right]
			right++
		} else {
			a[writer] = aux[left]
			left++
		}
	}
}

func sort(a, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, aux, lo, mid)
	sort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

func sortArray(nums []int) []int {
	aux := make([]int, len(nums))
	sort(nums, aux, 0, len(nums)-1)
	return nums
}
