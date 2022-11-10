package intsect2

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	freq := make(map[int]int)
	for _, n := range nums1 {
		freq[n]++
	}
	var k int
	for _, n := range nums2 {
		fq, ok := freq[n]
		if ok && fq > 0 {
			nums1[k] = n
			k++
			freq[n]--
		}
	}
	return nums1[:k]
}
