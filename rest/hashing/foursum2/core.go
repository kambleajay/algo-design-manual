package foursum2

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	freq := make(map[int]int)
	var ans int
	for _, a := range nums1 {
		for _, b := range nums2 {
			freq[a+b]++
		}
	}
	for _, c := range nums3 {
		for _, d := range nums4 {
			complement := -(c + d)
			ans += freq[complement]
		}
	}
	return ans
}
