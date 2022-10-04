package intsect

func intersection(nums1 []int, nums2 []int) []int {
	seen := make(map[int]bool)
	added := make(map[int]bool)
	for _, n := range nums1 {
		seen[n] = true
	}
	var result []int
	for _, n := range nums2 {
		if seen[n] && !added[n] {
			result = append(result, n)
			added[n] = true
		}
	}
	return result
}
