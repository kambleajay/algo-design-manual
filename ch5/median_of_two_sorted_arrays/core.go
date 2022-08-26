package median_of_two_sorted_arrays

func less(nums1 []int, nums2 []int, m int, n int, i int, j int) bool {
	return i < m && j < n && nums1[i] < nums2[j]
}

func hasNext(current, size int) bool {
	return current < size
}

func handleOddCase(nums1 []int, nums2 []int, m, n, size int) float64 {
	medianIndex := size / 2
	var i, j, count, answer int
	for i < size && j < size {
		if count == medianIndex {
			if less(nums1, nums2, m, n, i, j) || !hasNext(j, n) {
				answer = nums1[i]
			} else if j < n {
				answer = nums2[j]
			}
			break
		}
		if less(nums1, nums2, m, n, i, j) || !hasNext(j, n) {
			i++
		} else if j < n {
			j++
		}
		count++
	}
	return float64(answer)
}

func handleEvenCase(nums1 []int, nums2 []int, m, n, size int) float64 {
	index2 := size / 2
	index1 := index2 - 1
	var i, j, count, answer1, answer2 int
	for i < size && j < size {
		if count == index1 {
			if less(nums1, nums2, m, n, i, j) || !hasNext(j, n) {
				answer1 = nums1[i]
				i++
			} else if j < n {
				answer1 = nums2[j]
				j++
			}
			count++
			continue
		}
		if count == index2 {
			if less(nums1, nums2, m, n, i, j) || !hasNext(j, n) {
				answer2 = nums1[i]
			} else if j < n {
				answer2 = nums2[j]
			}
			break
		}
		if less(nums1, nums2, m, n, i, j) || !hasNext(j, n) {
			i++
		} else if j < n {
			j++
		}
		count++
	}
	return float64(answer1+answer2) / 2
}

func isEmpty(s []int) bool {
	return s == nil || len(s) == 0
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	if isEmpty(nums1) && len(nums2) == 1 {
		return float64(nums2[0])
	}
	if len(nums1) == 1 && isEmpty(nums2) {
		return float64(nums1[0])
	}
	m, n := len(nums1), len(nums2)
	size := m + n
	if size%2 != 0 {
		return handleOddCase(nums1, nums2, m, n, size)
	}
	return handleEvenCase(nums1, nums2, m, n, size)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays2(nums1, nums2)
}
