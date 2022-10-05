package intsect2

func intersect(nums1 []int, nums2 []int) []int {
	freq1 := make(map[int]int)
	for _, n := range nums1 {
		freq1[n]++
	}
	var result []int
	for _, n := range nums2 {
		fq1, ok1 := freq1[n]
		if ok1 && fq1 > 0 {
			result = append(result, n)
			freq1[n]--
		}
	}
	return result
}
