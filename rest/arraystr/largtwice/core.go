package largtwice

func dominantIndex(nums []int) int {
	var max1, max2, index int
	if nums[0] > nums[1] {
		max1, max2 = nums[1], nums[0]
		index = 0
	} else {
		max1, max2 = nums[0], nums[1]
		index = 1
	}
	for i, n := range nums[2:] {
		if n > max2 {
			max1 = max2
			max2 = n
			index = i + 2
		} else if n > max1 && n < max2 {
			max1 = n
		}
	}
	if max1*2 <= max2 {
		return index
	}
	return -1
}
