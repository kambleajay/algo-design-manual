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
	return kSum(nums, target, 4)
}

func kSum(nums []int, target, k int) [][]int {
	var res [][]int
	if nums == nil || len(nums) == 0 {
		return res
	}
	if nums[0] > target/k || nums[len(nums)-1] < target/k {
		return res
	}
	if k == 2 {
		return twoSum(nums, target)
	}
	for i, _ := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			resRest := kSum(nums[i+1:], target-nums[i], k-1)
			for _, rr := range resRest {
				res = append(res, append([]int{nums[i]}, rr...))
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	var res [][]int
	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum < target || (lo > 0 && nums[lo] == nums[lo-1]) {
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
