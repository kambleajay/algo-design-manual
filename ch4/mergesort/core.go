package mergesort

func less(a []int, i, j int) bool {
	return a[i] < a[j]
}

func merge(a, aux []int, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if less(aux, j, i) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func mergeSortIter(a, aux []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSortIter(a, aux, lo, mid)
	mergeSortIter(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

func mergeSort(a []int) {
	if a == nil || len(a) == 0 {
		return
	}
	aux := make([]int, len(a))
	mergeSortIter(a, aux, 0, len(a)-1)
}
