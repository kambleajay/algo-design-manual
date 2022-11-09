package kclosest

func findClosestElements(arr []int, k int, x int) []int {
	if k == len(arr) {
		return arr
	}
	lo, hi := 0, len(arr)-k
	for lo < hi {
		mid := lo + (hi-lo)/2
		if x-arr[mid] > arr[mid+k]-x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	var result []int
	for i := lo; i < lo+k; i++ {
		result = append(result, arr[i])
	}
	return result
}
