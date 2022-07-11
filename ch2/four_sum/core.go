package four_sum

import "sort"

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func fourSum(nums []int, target int) [][]int {
	if nums == nil || len(nums) < 4 {
		return [][]int{}
	}
	sort.Sort(IntSlice(nums))
	return kSum(nums, 0, 4, target)
}

func kSum(nums []int, start, k, target int) [][]int {
	var res [][]int
	if start == len(nums) {
		return res
	}
	avg := target / k
	if nums[start] > avg || nums[len(nums)-1] < avg {
		return res
	}
	if k == 2 {
		return twoSum(nums, start, target)
	}
	for i := start; i < len(nums); i++ {
		if i == start || nums[i] != nums[i-1] {
			for _, subset := range kSum(nums, i+1, k-1, target-nums[i]) {
				subset = append([]int{nums[i]}, subset...)
				res = append(res, subset)
			}
		}
	}
	return res
}

func twoSum(nums []int, start, target int) [][]int {
	var res [][]int
	lo, hi := start, len(nums)-1
	for hi > lo {
		sum := nums[lo] + nums[hi]
		if sum < target || (lo > start && nums[lo] == nums[lo-1]) {
			lo++
		} else if sum > target || (hi < len(nums)-1 && nums[hi] == nums[hi+1]) {
			hi--
		} else {
			res = append(res, []int{nums[lo], nums[hi]})
			lo++
			hi--
		}
	}
	return res
}
